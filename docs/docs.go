// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/projects": {
            "get": {
                "description": "Get a list of all projects associated with the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "List all projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.GetProject"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Create a new project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Project details",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Project created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error binding req object",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error executing db query",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a project by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Delete a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Project Deleted !!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No project found with the project ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/projects/{id}": {
            "get": {
                "description": "Get a project by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get a project by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetProject"
                        }
                    },
                    "404": {
                        "description": "No project found with the project ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a project by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Update a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Update project request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Project updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error binding req object",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error updating project",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sprints": {
            "get": {
                "description": "Lists down all sprints associated with the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprints"
                ],
                "summary": "List all sprints",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.GetSprint"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new sprint for a project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprints"
                ],
                "summary": "Create a new sprint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Sprint details",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateSprintRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sprint created successfully!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error executing db query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a sprint by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprints"
                ],
                "summary": "Delete a sprint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sprint ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sprint Deleted !!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error occurred while executing query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No sprint found with the sprintid",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sprints/{id}": {
            "get": {
                "description": "Get a sprint by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprints"
                ],
                "summary": "Get a sprint by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sprint ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.GetSprint"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Edit a sprint by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprints"
                ],
                "summary": "Edit a sprint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sprint ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Sprint details",
                        "name": "sprint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateSprintRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sprint updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error binding req object",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error updating sprint",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Get all tasks for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Fetch tasks",
                        "schema": {
                            "$ref": "#/definitions/common.Task"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a task with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Edit Task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tasks updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Title of the task",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Description of the task",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Type of the task (e.g., bug, feature)",
                        "name": "issue_type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Assignee of the task",
                        "name": "assignee",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID of the sprint the task belongs to",
                        "name": "sprint",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "ID of the project the task belongs to",
                        "name": "project_id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Points/effort estimation for the task",
                        "name": "points",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Reporter of the task",
                        "name": "reporter",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Additional comments for the task",
                        "name": "comments",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Status of the task",
                        "name": "status",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the task to delete",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tasks/list": {
            "get": {
                "description": "List all tasks for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "List tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.Task"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Task": {
            "type": "object",
            "properties": {
                "assignee": {
                    "type": "string"
                },
                "comments": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "issue_type": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                },
                "previousFields": {
                    "type": "object",
                    "additionalProperties": true
                },
                "project_id": {
                    "type": "integer"
                },
                "reporter": {
                    "type": "string"
                },
                "sprint_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "handlers.CreateProjectRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateSprintRequest": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "handlers.GetProject": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.GetSprint": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateProjectRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateSprintRequest": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tasks Service",
	Description:      "Tasks API in go using gin-framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
