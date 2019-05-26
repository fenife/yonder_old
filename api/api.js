let config = require('~/config')
let baseUrl = config.baseUrl

if (typeof window === 'undefined') {
  baseUrl = config.backendURl
}

const apiList = {
  // 获取文章列表
  getArticles: {
    url: baseUrl + "/articles",
    method: "GET"
  },
  // 获取文章详情
  getArticleDetail: {
    url: baseUrl + "/article/:id",
    method: "GET"
  },
  getCates: {
    url: baseUrl + "/categories",
    method: "GET"
  }
}

module.exports = apiList
