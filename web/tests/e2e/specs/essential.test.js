import { Selector } from 'testcafe'
import { testUrlBase } from '../testcafe.conf'

fixture(`Some essential features`)
  .page(`${testUrlBase}/`)

test(`Display available coupons`, async t => {
  const displayBlock = Selector('.available .coupon')
  await t
    .expect(displayBlock.exists).ok()
})

test(`Available coupons will update over time`, async t => {
  const beforeValue = await Selector('.available .coupon').innerText

  const requestButton = Selector('.request button')
  await t
    .setTestSpeed(0.01) // For demo purpose
    .click(requestButton)

  const afterValue = await Selector('.available .coupon').innerText
  await t
    .expect(afterValue).notEql(beforeValue)
})
