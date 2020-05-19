package rest

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/mesg-foundation/engine/execution"
	"github.com/mesg-foundation/engine/hash"
	"github.com/mesg-foundation/engine/x/execution/internal/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/execution/get/{hash}",
		queryGetHandlerFn(cliCtx),
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/execution/list",
		queryListHandlerFn(cliCtx),
	).Methods(http.MethodGet)
}

func queryGetHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryGet, vars["hash"])

		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryListHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryList)

		filter := types.ListFilter{}
		if param := r.FormValue("parentHash"); param != "" {
			h, err := hash.Decode(param)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			filter.ParentHash = h
		}

		if param := r.FormValue("eventHash"); param != "" {
			h, err := hash.Decode(param)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			filter.EventHash = h
		}

		if param := r.FormValue("instanceHash"); param != "" {
			h, err := hash.Decode(param)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			filter.InstanceHash = h
		}

		if param := r.FormValue("processHash"); param != "" {
			h, err := hash.Decode(param)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			filter.ProcessHash = h
		}

		if param := r.FormValue("status"); param != "" {
			status, ok := execution.Status_value[param]
			if !ok {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("%q is invalid", param))
				return
			}
			filter.Status = execution.Status(status)
		}

		if param := r.FormValue("nodeKey"); param != "" {
			filter.NodeKey = param
		}

		data, err := cliCtx.Codec.MarshalJSON(filter)
		if err != nil {
			return
		}
		res, height, err := cliCtx.QueryWithData(route, data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		execs := make([]*execution.Execution, 0)
		if err := cliCtx.Codec.UnmarshalJSON(res, &execs); err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		sort.Sort(sort.Reverse(execution.ByBlockHeight(execs)))

		start, end := client.Paginate(len(execs), page, limit, limit)
		if start < 0 || end < 0 {
			execs = []*execution.Execution{}
		} else {
			execs = execs[start:end]
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, execs)
	}
}
