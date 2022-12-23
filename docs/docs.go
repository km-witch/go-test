// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/block/data/{userid}": {
            "get": {
                "description": "유저의 블록 보유 확인 후 (없으면 생성 후)리턴",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "유저의 블록 보유 확인 및 생성",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write Block ID",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Resp_FindUserAndCreateBlock"
                        }
                    }
                }
            }
        },
        "/api/block/get/{userid}": {
            "get": {
                "description": "UserId 를 통해 Block정보를 가져옴.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "UserId -\u003e GetBlock",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write Block ID",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Block"
                        }
                    }
                }
            }
        },
        "/api/item/collection/{collectionid}": {
            "get": {
                "description": "Collection ID를 넣으면 Collection을 리턴함",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "Collection By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write Block ID",
                        "name": "collectionid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Collection"
                        }
                    }
                }
            }
        },
        "/api/item/group/{groupid}": {
            "get": {
                "description": "Product Group ID를 넣으면 -\u003e Group을 반환함.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "GetProductGroup By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write Block ID",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductGroup"
                        }
                    }
                }
            }
        },
        "/api/item/nft/{nftid}": {
            "get": {
                "description": "Nft ID를 넣으면 -\u003e NFT를 반환함.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "GetNftById By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write Nft ID",
                        "name": "nftid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Nft"
                        }
                    }
                }
            }
        },
        "/api/obj/airdrop": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "#에어드랍진행 (SaleID를 기준으로 트리이거나 또는 카드가 될 수 있음, 1인당 1개씩 수령가능)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "#에어드랍진행",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Plz Write",
                        "name": "ReqBody_Airdrop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqBody_Airdrop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Resp_Airdrop_Item"
                        }
                    }
                }
            }
        },
        "/api/obj/block/{blockid}": {
            "get": {
                "description": "#Obj 조회 By BlockID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "#Obj 조회 By BlockID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write Block ID",
                        "name": "blockid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Obj"
                            }
                        }
                    }
                }
            }
        },
        "/api/obj/msg": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "유저의 블록 보유 확인 후 (없으면 생성 후)리턴",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "obj message 작성",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Write User obj id and message",
                        "name": "ReqBody_ObjMessage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqBody_ObjMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Obj_msg"
                        }
                    }
                }
            }
        },
        "/api/obj/msg/count/{obj_id}": {
            "get": {
                "description": "Obj Message 갯수 조회 기능 (트리 크기 조정용)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "Obj Message 갯수 조회 기능 (트리 크기 조정용)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Object id",
                        "name": "obj_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Resp_GetMessageCount"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/api/obj/msg/del": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Obj msg의 is_active 값을 변경해 삭제처리 한다",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "Obj msg 삭제",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Write User obj id and obj_msg id",
                        "name": "ReqBody_ObjDel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqBody_ObjDel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Obj_msg"
                        }
                    }
                }
            }
        },
        "/api/obj/msg/paging/{page}/{limit}": {
            "get": {
                "description": "Obj messages 페이징 조회",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "Obj messages 페이징 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "페이지입력",
                        "name": "page",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "조회갯수제한",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Response_ReadObjMessages"
                        }
                    }
                }
            }
        },
        "/api/obj/msg/{id}": {
            "get": {
                "description": "Obj Msg 단일 조회 기능 (단일, ID값 기준)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ready"
                ],
                "summary": "Obj Msg 단일 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "오브제 메세지 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Obj_msg"
                        }
                    }
                }
            }
        },
        "/api/obj/userid": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "#JWT Token을 헤더에 포함하면 Obj를 조회함",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "#JWT Token을 헤더에 포함하면 Obj를 조회함",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Resp_GetObjsByUserId"
                        }
                    }
                }
            }
        },
        "/api/user/": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "유저의 블록 보유 확인 후 (없으면 생성 후)리턴",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Main"
                ],
                "summary": "#유저 접속시 호출 초기화 및 유저 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Write User Token",
                        "name": "ReqBody_Token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqBody_Token"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Resp_FindUserBlockData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ReqBody_Airdrop": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "msg_role": {
                    "type": "integer"
                },
                "pos": {
                    "type": "string"
                },
                "rot": {
                    "type": "string"
                },
                "sale_id": {
                    "description": "얘로 Tree면 2번을, 카드면 3번을 넣어주세요.",
                    "type": "string"
                }
            }
        },
        "controller.ReqBody_ObjDel": {
            "type": "object",
            "properties": {
                "objId": {
                    "type": "integer"
                },
                "objMsgId": {
                    "type": "integer"
                }
            }
        },
        "controller.ReqBody_ObjMessage": {
            "type": "object",
            "properties": {
                "objId": {
                    "type": "integer"
                },
                "objMessage": {
                    "type": "string"
                },
                "userNickname": {
                    "type": "string"
                }
            }
        },
        "controller.ReqBody_Token": {
            "type": "object",
            "properties": {
                "nickName": {
                    "type": "string"
                }
            }
        },
        "controller.Resp_Airdrop_Item": {
            "type": "object",
            "properties": {
                "payload": {
                    "$ref": "#/definitions/model.Obj"
                }
            }
        },
        "controller.Resp_FindUserAndCreateBlock": {
            "type": "object",
            "properties": {
                "block": {
                    "$ref": "#/definitions/model.Block"
                },
                "objs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Obj"
                    }
                }
            }
        },
        "controller.Resp_FindUserBlockData": {
            "type": "object",
            "properties": {
                "block": {
                    "$ref": "#/definitions/model.Block"
                },
                "objs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Obj_with_productid"
                    }
                }
            }
        },
        "controller.Resp_GetMessageCount": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "integer",
                    "example": 24
                }
            }
        },
        "controller.Resp_GetObjsByUserId": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Obj"
                    }
                }
            }
        },
        "controller.Response_ReadObjMessages": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Obj_msg"
                    }
                }
            }
        },
        "model.Block": {
            "type": "object",
            "properties": {
                "created_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "thema": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Collection": {
            "type": "object",
            "required": [
                "name_en",
                "name_ko",
                "publisher_en",
                "publisher_ko"
            ],
            "properties": {
                "created_time": {
                    "type": "string"
                },
                "discord_url": {
                    "type": "string"
                },
                "id": {
                    "description": "PK",
                    "type": "integer"
                },
                "name_en": {
                    "type": "string"
                },
                "name_ko": {
                    "type": "string"
                },
                "opensea_url": {
                    "type": "string"
                },
                "publisher_en": {
                    "type": "string"
                },
                "publisher_ko": {
                    "type": "string"
                },
                "twitter_url": {
                    "type": "string"
                },
                "ww_url": {
                    "type": "string"
                }
            }
        },
        "model.Nft": {
            "type": "object",
            "required": [
                "description",
                "name_en",
                "name_ko",
                "product_id",
                "token_id"
            ],
            "properties": {
                "created_time": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name_en": {
                    "type": "string"
                },
                "name_ko": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "properties": {
                    "type": "string"
                },
                "token_id": {
                    "type": "integer"
                },
                "updated_time": {
                    "type": "string"
                },
                "wallet_id": {
                    "type": "integer"
                }
            }
        },
        "model.Obj": {
            "type": "object",
            "required": [
                "block_id",
                "pos",
                "rot",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "block_id": {
                    "type": "integer"
                },
                "building_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "description": "PK",
                    "type": "integer"
                },
                "msg_role": {
                    "type": "integer"
                },
                "nft_id": {
                    "type": "integer"
                },
                "pos": {
                    "type": "string"
                },
                "rot": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_user": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Obj_msg": {
            "type": "object",
            "required": [
                "created_user",
                "message",
                "obj_id",
                "updated_user",
                "user_nickname"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_user": {
                    "type": "integer"
                },
                "id": {
                    "description": "PK",
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "obj_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_user": {
                    "type": "integer"
                },
                "user_nickname": {
                    "type": "string"
                }
            }
        },
        "model.Obj_with_productid": {
            "type": "object",
            "required": [
                "block_id",
                "pos",
                "rot",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "block_id": {
                    "type": "integer"
                },
                "building_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "description": "PK",
                    "type": "integer"
                },
                "msg_role": {
                    "type": "integer"
                },
                "nft_id": {
                    "type": "integer"
                },
                "pos": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "rot": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_user": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.ProductGroup": {
            "type": "object",
            "required": [
                "collection_id",
                "contract"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "collection_id": {
                    "type": "integer"
                },
                "contract": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "message_amount": {
                    "type": "integer"
                },
                "message_role": {
                    "description": "MSG Role을 여기서 주어야 하는가..에 대한 의문.",
                    "type": "string"
                },
                "metadata": {
                    "type": "string"
                },
                "name_en": {
                    "type": "string"
                },
                "name_ko": {
                    "type": "string"
                },
                "properties": {
                    "type": "string"
                },
                "snap": {
                    "type": "string"
                },
                "thumbnail_url": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
