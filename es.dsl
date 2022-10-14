GET _search
{
  "query": {
    "match_all": {}
  }
}

DELETE /test

PUT /test
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  }
}

DELETE /employee

#非结构化新建索引
PUT /employee/_doc/4
{
  "name":"凯杰23",
  "age":35
}

PUT /employee/_doc/1
{
  "name":"凯杰123"
}

GET /employee/_doc/1

POST /employee/_update/1
{
  "doc":{
    "name":"凯"
  }
}

POST /employee/_create/2
{
  "name":"123",
  "age":"30"
}

#删除文档
DELETE /employee/_doc/2

#查询全部文档
GET /employee/_search

#使用结构化方式创建索引
PUT /employee
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "name":{"type": "text"},
      "age":{"type": "integer"}
    }
  }
}

#不带条件查询
GET /employee/_search
{
  "query":{
    "match_all": {}
  }
}

#分页查询
GET /employee/_search
{
  "query":{
    "match_all": {}
  },
  "from":0,
  "size":2
}

#带关键字查询
GET /employee/_search
{
  "query": {
    "match": {
      "name": "凯"
    }
  }
}

#带排序
GET /employee/_search
{
  "query": {
    "match": {
      "name": "凯"
    }
  },
  "sort":[
    {"age":{"order":"desc"}}
  ]
}

#带filter
GET /employee/_search
{
  "query": {
    "bool": {
      "filter":[
        {"term":{"name":"凯4234324"}}
      ]
    }
  }
}

#带聚合
GET /employee/_search
{
  "query": {
    "match": {
      "name": "凯"
    }
  },
  "sort":[
    {"age":{"order":"desc"}}
  ],
  "aggs": {
    "group_by_age": {
      "terms": {
        "field": "age"
      }
    }
  }
}

#新建索引
PUT /movie/_doc/1
{
  "name": "Eating an apple a day & keeps the doctor away"
}

GET /movie/_search
{
  "query": {
    "match": {
      "name": "apples"
    }
  }
}

#使用analyze api查看分词状态
GET /movie/_analyze
{
  "field": "name",
  "text": "Eating an apple a day & keeps the doctor away"
}

#使用analyze api查看分词状态
GET /movie/_analyze
{
  "field": "name",
  "text": "apples"
}

DELETE /movie

#使用结构化方式重新创建索引
PUT /movie
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "name":{"type": "text", "analyzer": "english"}
    }
  }
}

GET /movie/_search

#使用analyze api查看分词状态
GET /movie/_analyze
{
  "field": "name",
  "text": "Eating an apples a day & keeps the doctor away"
}