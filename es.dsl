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

#开始玩转tmdb
PUT /movie
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "title":{
        "type": "text",
        "analyzer": "english"
      },
      "tagline":{
        "type": "text",
        "analyzer": "english"
      },
      "release_date":{
          "type":"date",
          "format":"8yyyy/MM/dd||yyyy/M/dd||yyyy/MM/d||yyyy/M/d"
      },
      "popularity":{
        "type":"double"
      },
      "overview":{
        "type":"text",
        "analyzer":"english"
      },
      "cast":{
        "type": "object",
        "properties": {
          "character":{
            "type":"text",
            "analyzer":"standard"
          },
          "name":{
            "type":"text",
            "analyzer":"standard"
          }
        }
      }
    }
  }
}

#搜索内容 match,分词后被索引查询
GET /movie/_search
{
  "query": {
    "match": {
      "title": "steve zissou"
    }
  }
}

GET /movie/_analyze
{
  "field": "title",
  "text": "steve zissou"
}

#搜索内容 term，不进行分词分析，直接查询
GET /movie/_search
{
  "query": {
    "term": {
      "title": "steve zissou"
    }
  }
}

#分词后的and和or逻辑
GET /movie/_search
{
  "query": {
    "match": {
      "title": "basketball with cartoom aliens"
    }
  }
}

#改成and
GET /movie/_search
{
  "query": {
    "match": {
      "title": {
        "query": "basketball with cartoom aliens",
        "operator": "and"
      }
    }
  }
}

#最小词匹配项
GET /movie/_search
{
  "query": {
    "match": {
      "title": {
        "query": "basketball love aliens",
        "operator": "or",
        "minimum_should_match": 2
      }
    }
  }
}

#短语查询
GET /movie/_search
{
  "query": {
    "match_phrase": {
      "title": "steve zissou"
    }
  }
}

#多字段查询
GET /movie/_search
{
  "explain": true,
  "query": {
    "multi_match": {
      "query": "basketball with cartoom aliens",
      "fields": ["title","overview"]
    }
  }
}

GET /movie/_search
{
  "explain": true,
  "query": {
    "match": {
      "title": "steve"
    }
  }
}

#优化多字段查询
GET /movie/_search
{
  "explain": true,
  "query": {
    "multi_match": {
      "query": "basketball with cartoom aliens",
      "fields": ["title^10","overview"],
      "tie_breaker": 0.3
    }
  }
}

#bool查询
#must:必须都为true
#must not:必须都是false
#shoule:其中有一个为true即可
#为true越多则分越高
GET /movie/_search
{
  "explain": true,
  "query": {
    "bool": {
      "should": [
        {"match":{"title": "basketball with cartoom aliens"}},
        {"match":{"overview": "basketball with cartoom aliens"}}
      ]
    }
  }
}

#不同的multi_query其实是有不同的type
#best_fields:默认的得分方式，取得最高的分数作为对应文档的分数
GET /movie/_search
{
  "explain": true,
  "query": {
    "multi_match": {
      "query": "basketball with cartoom aliens",
      "fields": ["title", "overview"],
      "type": "best_fields"
    }
  }
}

GET /movie/_search
{
  "explain": true,
  "query": {
    "dis_max": {
      "queries": [
        {"match":{"title": "basketball with cartoom aliens"}},
        {"match":{"overview": "basketball with cartoom aliens"}}
      ]
    }
  }
}

GET /movie/_validate/query?explain
{
  "query": {
    "multi_match": {
      "query": "basketball with cartoom aliens",
      "fields": ["title^10", "overview"],
      "type": "best_fields"
    }
  }
}

#most_fields:考虑绝大多数（所有的）
GET /movie/_search
{
  "explain": true,
  "query": {
    "multi_match": {
      "query": "basketball with cartoom aliens",
      "fields": ["title^10", "overview^0.1"],
      "type": "most_fields"
    }
  }
}

GET /movie/_validate/query?explain
{
  "query": {
    "multi_match": {
      "query": "basketball with cartoom aliens",
      "fields": ["title^10", "overview^0.1"],
      "type": "most_fields"
    }
  }
}

#cross_fields:以分词为单位计算栏位的总分，词导向
GET /movie/_search
{
  "explain": true,
  "query": {
    "multi_match": {
      "query": "steve jobs",
      "fields": ["title", "overview"],
      "type": "cross_fields"
    }
  }
}

GET /movie/_validate/query?explain
{
  "query": {
    "multi_match": {
      "query": "steve jobs",
      "fields": ["title", "overview"],
      "type": "cross_fields"
    }
  }
}

#query string
#便于利用AND OR NOT
GET /movie/_search
{
  "query": {
    "query_string": {
      "fields": ["title"],
      "query": "steve OR jobs"
    }
  }
}

#filter过滤查询
#单条件过滤
GET /movie/_search
{
  "query": {
    "bool": {
      "filter": {
        "term":{"title":"steve"}
      }
    }
  }
}

#多条件过滤
GET /movie/_search
{
  "query": {
    "bool": {
      "filter": [
        {"term":{"title": "steve"}},
        {"term": {"cast.name": "gaspard"}},
        {"range":{"release_date": {"lte": "2015/01/01"}}},
        {"range":{"popularity": {"gte":"25"}}}
      ]
    }
  },
  "sort": [
    {
      "popularity": {
        "order": "desc"
      }
    }
  ]
}

#带match打分的filter
GET /movie/_search
{
  "query": {
    "bool": {
      "should": [
        {"match": {
          "title": "life"
        }}
      ],
      "filter": [
        {"term":{"title": "steve"}},
        {"term": {"cast.name": "gaspard"}},
        {"range":{"release_date": {"lte": "2015/01/01"}}},
        {"range":{"popularity": {"gte":"25"}}}
      ]
    }
  }
}

#functionscore
GET /movie/_search
{
  "explain": true,
  "query": {
    "function_score": {
        "query": {
          "multi_match": {
          "query": "steve jobs",
          "fields": ["title", "overview"],
          "operator": "or",
          "type": "most_fields"
        }
      },
      "functions": [
        {
          "field_value_factor":{
            "field":"popularity",
            "modifier":"log2p",
            "factor":10
          }
        },
        {
          "field_value_factor":{
            "field":"popularity",
            "modifier":"log2p",
            "factor":5
          }
        }
      ],
      "score_mode": "sum",
      "boost_mode": "sum"
    }
  }
}
