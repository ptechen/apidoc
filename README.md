# apidoc

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ptechen/apidoc)
[![Build Status](https://travis-ci.com/ptechen/apidoc.svg?branch=master)](https://travis-ci.com/ptechen/apidoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/ptechen/apidoc)](https://goreportcard.com/report/github.com/ptechen/apidoc)
[![codecov](https://codecov.io/gh/ptechen/apidoc/branch/master/graph/badge.svg)](https://codecov.io/gh/ptechen/apidoc)

自动生成 apidoc 文档

example:
        
        type ReqParams struct {
        	Id    int    `json:"id" comment:"test"`
        	Name  string `json:"name" comment:"test"`
        	inner string
        	Age   int `json:"age" comment:"test"`
        	Xi    struct {
        		Name string `json:"name" comment:"test"`
        	} `json:"xi" comment:"test"`
        }
        
        type ResParams struct {
        	Id    int    `json:"id" comment:"test"`
        	Name  string `json:"name" comment:"test"`
        	inner string
        	Age   int `json:"age" comment:"test"`
        	Xi    struct {
        		Name string `json:"name" comment:"test"`
        		Info struct {
        			Age int `json:"age" comment:"age"`
        		} `json:"info" comment:"user_info"`
        	} `json:"xi" comment:"test"`
        }
        
        req := ReqParams{}
    	res := ResParams{}
    	f, _ := os.Create("./apidoc/apidoc")
    	api := &Api{
    		Output:   f,
    		Method:   "post",
    		Route:    "/test",
    		Desc:     "123",
    		Version:  "v0.0.1",
    		ApiName:  "user",
    		ApiGroup: "user",
    	}
    	api.ApiDoc(req, res)
   
result:

    /**
     * @api { post } /test   123
     * @apiVersion v0.0.1
     * @apiName user
     * @apiGroup user 
     * @apiParam { Number } 	 id 	 test 
     * @apiParam { String } 	 name 	 test 
     * @apiParam { Number } 	 age 	 test 
     * @apiParam { Object } 	 xi 	 test 
     * @apiParam { String } 	 xi.name 	 test 
     * @apiParamExample {json} Request-Example:
     * { 
     *     "id": 0, 
     *     "name": "", 
     *     "age": 0, 
     *     "xi": { 
     *         "name": "" 
     *     } 
     * } 
     * 
     * @apiSuccess { Number } 	 id 	 test
     * @apiSuccess { String } 	 name 	 test
     * @apiSuccess { Number } 	 age 	 test
     * @apiSuccess { Object } 	 xi 	 test
     * @apiSuccess { String } 	 xi.name 	 test
     * @apiSuccess { Object } 	 xi.info 	 user_info
     * @apiSuccess { Number } 	 xi.info.age 	 age
     * @apiSuccessExample {json} Request-Example:
     * { 
     *     "id": 0, 
     *     "name": "", 
     *     "age": 0, 
     *     "xi": { 
     *         "name": "", 
     *         "info": { 
     *             "age": 0 
     *         } 
     *     } 
     * } 
     */
    
