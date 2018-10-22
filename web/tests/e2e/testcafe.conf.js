const fs = require('fs')
const path = require('path')

const srcPaths = fs.readdirSync(path.resolve(__dirname, 'specs/'))
  .map(v => path.resolve(__dirname, `./specs/${v}`))

module.exports = {
  src: srcPaths,
  host: 'localhost',
  testUrlBase: 'http://localhost:8080',
  browsers: 'chrome'
}
