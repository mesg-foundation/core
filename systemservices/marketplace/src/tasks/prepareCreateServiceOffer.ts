import { TaskInputs, TaskOutputs } from "mesg-js/lib/service"
import { Marketplace } from "../contracts/Marketplace"
import { toUnit, stringToHex, CreateTransaction } from "../contracts/utils";
import BigNumber from "bignumber.js";
import { getService } from "../contracts/service";
import * as assert from "assert";
import { getServiceVersionCount } from "../contracts/version";

export default (
  marketplace: Marketplace,
  createTransaction: CreateTransaction
) => async (inputs: TaskInputs, outputs: TaskOutputs): Promise<void> => {
  try {
    // check inputs
    const sid = inputs.sid
    const duration = new BigNumber(inputs.duration)
    assert.ok(duration.isPositive(), 'duration must be strictly positive')

    // check service
    const service = await getService(marketplace, sid)

    // check ownership
    assert.strictEqual(inputs.from.toLowerCase(), service.owner.toLowerCase(), `service's owner is different`)

    // check service version
    const versionsLength = await getServiceVersionCount(marketplace, sid)
    assert.ok(!versionsLength.isEqualTo(0), 'cannot create an offer on a service with no version')

    // create transaction
    const transactionData = marketplace.methods.createServiceOffer(
      stringToHex(sid),
      toUnit(inputs.price),
      duration.toString()
    ).encodeABI()
    return outputs.success(await createTransaction(marketplace, inputs, transactionData))
  }
  catch (error) {
    console.error('error in prepareCreateServiceOffer', error)
    return outputs.error({ message: error.message })
  }
}
