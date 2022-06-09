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
        "/portal/ban_user": {
            "post": {
                "description": "BanUser",
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"禁言成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/check_noob": {
            "post": {
                "description": "check user is noob",
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/get_banned_users": {
            "post": {
                "description": "GetBannedUsers",
                "tags": [
                    "Portal"
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"users\": users}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/get_green": {
            "post": {
                "description": "get user green status",
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"data\": greenStatus}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/get_greenbirds": {
            "post": {
                "description": "GetGreenbirds",
                "tags": [
                    "Portal"
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"data\": data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/get_hot_posts": {
            "post": {
                "description": "get posts which get the most highest views",
                "tags": [
                    "Portal"
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"data\": data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/get_notifications": {
            "post": {
                "description": "Get user notifications",
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "length",
                        "name": "length",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"data\": notifications}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/get_user_message": {
            "post": {
                "description": "Get user system message",
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "length",
                        "name": "length",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"获取成功\", \"data\": SysMessages}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/save_greenbirds": {
            "post": {
                "description": "SaveGreenbird",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "description": "新手上路信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GreenbirdData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"保存成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/upload_file": {
            "post": {
                "description": "UploadImage",
                "tags": [
                    "Portal"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "图片",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"上传成功\", \"url\": url}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/add_post_tag": {
            "post": {
                "description": "Add post a tag",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"添加标签成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/comment/create": {
            "post": {
                "description": "Create a comment",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户评论成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/comment/delete": {
            "post": {
                "description": "Delete a comment",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "comment_id",
                        "name": "comment_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户评论成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/comment/like": {
            "post": {
                "description": "Like a comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "LikeCommentData",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LikeCommentData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"点赞成功\", \"commentlike\": commentlike}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/create": {
            "post": {
                "description": "Create a post --note-- section in [0,3]",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "22",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePostData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"发布成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/delete": {
            "post": {
                "description": "delete post and sub user exp",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"删除文章成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get": {
            "post": {
                "description": "Get posts with offset and length",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "22",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetPostsData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取文章成功\", \"data\": data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_post_comments": {
            "post": {
                "description": "Get post comments",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "length",
                        "name": "length",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取评论成功\", \"comments\":comments}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_post_tags": {
            "post": {
                "description": "Get post tags",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "post_id",
                        "name": "post_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取标签成功\", \"post_tags\":postTags}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_section_tags": {
            "get": {
                "description": "Get section tags",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "section",
                        "name": "section",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取标签成功\", \"tags\":tags}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/get_user_posts": {
            "post": {
                "description": "Get user posts",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "length",
                        "name": "length",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取文章成功\", \"data\": data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/like": {
            "post": {
                "description": "Like a post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "LikePostData",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LikePostData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"点赞成功\", \"postlike\": postlike}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/read": {
            "post": {
                "description": "read post and add user exp",
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取文章成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/search": {
            "post": {
                "description": "Search post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "22",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SearchPostsData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"搜索成功\", \"data\": data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "post": {
                "description": "ShowUserInfo",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"查询成功\", \"data\": user}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"登录成功\",\"data\": user}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password1",
                        "name": "password1",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password2",
                        "name": "password2",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update_exp": {
            "post": {
                "description": "update user exp",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "exp",
                        "name": "exp",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update_info": {
            "post": {
                "description": "UpdateInfo",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sex",
                        "name": "sex",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "age",
                        "name": "age",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update_password": {
            "post": {
                "description": "UpdatePassword",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "old_password",
                        "name": "old_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password1",
                        "name": "password1",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password2",
                        "name": "password2",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/upload_avatar": {
            "post": {
                "description": "upload user avatar",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "avatar",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": true, \"message\": \"上传成功\", \"avatar_url\": url}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreatePostData": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "section": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.GetPostsData": {
            "type": "object",
            "properties": {
                "length": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "order": {
                    "type": "string"
                },
                "section": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Greenbird": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.GreenbirdData": {
            "type": "object",
            "properties": {
                "greenBirds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Greenbird"
                    }
                }
            }
        },
        "model.LikeCommentData": {
            "type": "object",
            "properties": {
                "comment_id": {
                    "type": "integer"
                },
                "like_or_dislike": {
                    "type": "boolean"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.LikePostData": {
            "type": "object",
            "properties": {
                "like_or_dislike": {
                    "type": "boolean"
                },
                "post_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.SearchPostsData": {
            "type": "object",
            "properties": {
                "fliters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "length": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "order": {
                    "type": "string"
                },
                "section": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "name": {
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
