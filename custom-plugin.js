const {analyzeCommits} = require("semantic-release-gitmoji");

async function analyzeCommit(pluginConfig, context) {
  console.log("Hello World !!!!")
  return analyzeCommits(pluginConfig, context)
}

module.exports = {
  analyzeCommit
}
