
# Task Manager API Documentation

This document describes the endpoints for the Task Manager REST API.

---

## 1. GET /tasks

**Description:**  
Fetch a list of all tasks.

**Request:**  
```http
GET {{base_url}}/tasks
````

**Example cURL:**

```bash
curl --location --globoff '{{base_url}}/tasks'
```

**Response:**

```json
{
  "tasks": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "Description for Task 1",
      "due_date": "2025-07-17T23:24:04.5309372+03:00",
      "status": "Pending"
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "Description for Task 2",
      "due_date": "2025-07-18T23:24:04.5309372+03:00",
      "status": "In Progress"
    },
    {
      "id": "3",
      "title": "Task 3",
      "description": "Description for Task 3",
      "due_date": "2025-07-19T23:24:04.5309372+03:00",
      "status": "Completed"
    }
  ]
}
```

---

## 2. GET /tasks/\:id

**Description:**
Fetch a task by its ID.

**Request:**

```http
GET {{base_url}}/tasks/2
```

**Example cURL:**

```bash
curl --location --globoff '{{base_url}}/tasks/2'
```

**Response:**

```json
{
  "id": "2",
  "title": "Task 2",
  "description": "Description for Task 2",
  "due_date": "2025-07-18T23:24:04.5309372+03:00",
  "status": "In Progress"
}
```

---

## 3. POST /tasks

**Description:**
Create a new task.

**Request:**

```http
POST {{base_url}}/tasks
```

**Request Body:**

```json
{
  "title": "Updated Task Title",
  "description": "This is the updated description.",
  "due_date": "2025-07-25T15:00:00Z",
  "status": "Completed"
}
```

**Example cURL:**

```bash
curl --location --globoff '{{base_url}}/tasks' \
--data '{
  "title": "Updated Task Title",
  "description": "This is the updated description.",
  "due_date": "2025-07-25T15:00:00Z",
  "status": "Completed"
}'
```

**Response:**

```json
{
  "message": "Task created successfully",
  "task": {
    "id": "",
    "title": "Updated Task Title",
    "description": "This is the updated description.",
    "due_date": "2025-07-25T15:00:00Z",
    "status": "Completed"
  }
}
```

---

## 4. PUT /tasks/\:id

**Description:**
Update an existing task.

**Request:**

```http
PUT {{base_url}}/tasks/3
```

**Request Body:**

```json
{
  "status": "In Progress"
}
```

**Example cURL:**

```bash
curl --location --globoff --request PUT '{{base_url}}/tasks/3' \
--data '{
  "status": "In Progress"
}'
```

**Response:**

```json
{
  "message": "Task updated successfully",
  "task": {
    "id": "3",
    "title": "Task 3",
    "description": "Description for Task 3",
    "due_date": "2025-07-19T23:24:04.5309372+03:00",
    "status": "In Progress"
  }
}

## 5. DELETE /tasks/\:id

**Description:**
Delete a task by its ID.

**Request:**

```http
DELETE {{base_url}}/tasks/3
```

**Example cURL:**

```bash
curl --location --globoff --request DELETE '{{base_url}}/tasks/3'
```

**Response:**

```json
{
  "message": "Task deleted successfully"
}
```

---

# Notes

* Replace `{{base_url}}` with your actual server URL, e.g., `http://localhost:8080`
* Replace `:id` in URLs with the actual task ID
* All timestamps use ISO 8601 format

