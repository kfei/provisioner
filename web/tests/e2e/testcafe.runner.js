process.env.NODE_ENV = 'testing'
const chalk = require('chalk')
const createTestCafe = require('testcafe')
const {
  src,
  host,
  browsers
} = require('./testcafe.conf')

let testcafe

createTestCafe(host)
  .then(tc => {
    testcafe = tc
    return testcafe.createRunner()
      .src(src)
      .browsers(browsers)
      .run()
  })
  .then(count => {
    if (count > 0) {
      console.log(chalk.red(`  e2e test failed: ${count}\n`))
    } else {
      console.log(chalk.cyan('  e2e test all passed.\n'))
    }
    testcafe.close()
    process.exit()
  })
