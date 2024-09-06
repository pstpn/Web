// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/employee-photo": {
            "get": {
                "description": "Метод для получения фотографии своего профиля",
                "tags": [
                    "employee"
                ],
                "summary": "Получение фотографии своего профиля",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Проверка на жизнеспособность",
                "tags": [
                    "system"
                ],
                "summary": "Проверка здоровья",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/infocard-photos/{id}": {
            "get": {
                "description": "Метод для получения элемента коллекции фотографий сотрудников",
                "tags": [
                    "admin"
                ],
                "summary": "Получение элемента коллекции фотографий сотрудников",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/infocards": {
            "get": {
                "description": "Метод для получения коллекции информационных карточек",
                "tags": [
                    "admin"
                ],
                "summary": "Получение коллекции информационных карточек",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/infocards/{id}": {
            "get": {
                "description": "Метод для получения элемента коллекции информационных карточек",
                "tags": [
                    "admin"
                ],
                "summary": "Получение элемента коллекции информационных карточек",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            },
            "patch": {
                "description": "Метод для подтверждения информационной карточки сотрудника",
                "tags": [
                    "admin"
                ],
                "summary": "Подтверждение информационной карточки сотрудника",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Метод для авторизации пользователя",
                "tags": [
                    "auth"
                ],
                "summary": "Авторизация пользователя",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/passages": {
            "post": {
                "description": "Метод для записи о проходе через КПП",
                "tags": [
                    "admin"
                ],
                "summary": "Запись информации о проходе через КПП",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "description": "Метод для получения профиля",
                "tags": [
                    "employee"
                ],
                "summary": "Получение профиля",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            },
            "post": {
                "description": "Метод для заполнения профиля",
                "tags": [
                    "employee"
                ],
                "summary": "Заполнение профиля",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Метод для обновления токенов доступа пользователя",
                "tags": [
                    "auth"
                ],
                "summary": "Обновление токенов доступа",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Метод для регистрации пользователя",
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "responses": {
                    "200": {
                        "description": "Сервис жив",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Сервис мертв"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Идентификация на КПП",
	Description:      "В основе лежит мой курсовой проект 😎",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
