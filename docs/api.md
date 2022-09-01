# TIDIFY-API

### Docs History

| 버전 | 변경일     | 변경 내용 | 작성자                        |
| ---- | ---------- | --------- | ----------------------------- |
| v0.1.0 | 2022.08.24 | 형식 작성 | 박재용(scorpion@dgu.ac.k) |
| v0.2.0 | 2022.08.26 | 응답코드 세분화 | 박재용(scorpion@dgu.ac.k) |
| v0.3.0 | 2022.08.28 | 응답코드 세분화 | 박재용(scorpion@dgu.ac.k) |
| v0.4.0 | 2022.08.30 | 기능 추가 | 박재용(scorpion@dgu.ac.k) |
| v0.5.0 | 2022.08.31 | 기능 추가 | 박재용(scorpion@dgu.ac.k) |
| v1.0.0 | 2022.08.31 | 문서 배포 | 박재용(scorpion@dgu.ac.k) |



### Server 자체 응답코드 (api_response.return_code)
| Code | Message     | HTTP Status Code | Description |
| ---- | ---------- | --------- | ----------------------------- |
| N200 | OK | 200 | 정상적으로 완료됨 |
| E300 | Token authentication error | 401 | 토큰 인증 에러 |
| E301 | Already expired token | 401 | 토큰이 만료됨 |
| E302 | No permission | 401 | 권한이 없음 |
| E303 | Internal Server Error Occured | 500 | 내부 서버 에러 |
| E311 | Some of request data is empty | 400 | 일부 요청 데이터가 누락됨 |
| E312 | Please check format of request queries | 400 | 일부 요청 쿼리 파라미터가 비정상임 |
| E313 | Please check request datas | 400 | 일부 요청 데이터가 부적합함 |
| E320 | Can't communicate with internal database | 500 | 데이터 베이스 에러 |
|      |            |           |                               |


## 인증 API

### API 목록

| No   | Method | URI    | 설명  | 권한  | 
| ---- | ------ | ------ | ----- | ----- |
| 1    | GET   | /auth/google | 구글 계정 회원가입 및 로그인 | 없음 |
| 2    | GET   | /auth/kakao | 카카오 계정 회원가입 및 로그인 | 없음 |
| 3    | POST   | /auth/apple | 애플 계정 회원가입 및 로그인 | 없음 |
| 4    | GET   | /signin | 토큰 갱신(※refresh-token cookie 필요) | 사용자 |

### 1.  GET /auth/google
구글 계정으로 로그인합니다. 구글 로그인 페이지로 Redirect 된 후, 인증을 완료하면 로그인되어 Access Token과 Refresh Token이 발급됩니다.
<br>※단, 최초 로그인 시 자동으로 가입됨

#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | Cookie: access-token, refresh-token 발급 완료 | 

### 2.  GET /auth/kakao
카카오 계정으로 로그인합니다. 카카오 로그인 페이지로 Redirect 된 후, 인증을 완료하면 로그인되어 Access Token과 Refresh Token이 발급됩니다.
<br>※단, 최초 로그인 시 자동으로 가입됨

#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | Cookie: access-token, refresh-token 발급 완료 | 

### 3.  POST /auth/apple
애플 계정으로 로그인합니다. Request Body에 Identity Token을 담아서 요청하면 로그인되어 Access Token과 Refresh Token이 발급됩니다.
<br>※단, 최초 로그인 시 자동으로 가입됨

#### Request Body
JSON 형식

| 항목        | 타입    | 설명                 |
| ----------- | ------- | -------------------- |
| id_token | String | Apple에서 발급받은 Identity Token |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | Cookie: access-token, refresh-token 발급 완료 | 

### 4.  GET /signin
Access Token이 만료되면 Refresh Token을 이용하여 Access Token을 재발급 받을 수 있습니다.
<br>※단, Refresh Token도 만료된 경우 재로그인 필요.

#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| refresh-token | Server로부터 발급받은 만료되지 않은 토큰 |

#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | Cookie: access-token 재발급 완료 | 
| E300   | Cookie: refresh-token 만료로 재발급 불가 | 

#### 요청 및 응답 예시
[요청]
```
GET http://localhost:8888/signin
```
[Cookie]
```
refresh-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjI1Mzc2MjV9.meC-WAKjTGLAVJ2uXVsMCZGRA2c4W0SdTMr02QRnokY; Path=/; HttpOnly; Expires=Wed, 07 Sep 2022 08:00:25 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NjQwNzR9.Xs07phCjB8cqNTTrwyUprjX9Rwevp08cxZfJpURHeOU; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 14:41:14 GMT;
```

## 북마크 API

### API 목록

| No   | Method | URI    | 설명  | 권한  | 
| ---- | ------ | ------ | ----- | ----- |
| 1    | GET   | /bookmarks | 본인 북마크 조회 | 사용자 |
| 2    | POST   | /bookmarks | 본인 북마크 생성 | 사용자 |
| 3    | DELETE   | /bookmarks | 본인 북마크 삭제 | 사용자 |
| 4    | PUT   | /bookmarks | 본인 북마크 갱신 | 사용자 |

### 1.  GET /bookmarks
자신의 북마크 리스트를 조건에 맞게 불러온다.
#### Request Query
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| start   | Integer | N(default: 0)   | 조회 시작 순번 |
| count   | Integer | Y(0보다 큰 수)   | 조회 개수 |
| folder   | Integer | N(default: 0 - 전체 조회)   | 폴더 ID로 특정 폴더 조회 |
| keyword   | String | N   | 북마크 URL 또는 제목으로 리스트 조회 |

#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 북마크 리스트 불러오기 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E312   | 잘못된 형식의 Request Query | 


#### 요청 및 응답 예시
[요청]
```
GET http://localhost:8888/bookmarks?start=0&count=5&folder=0
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "list": [
        {
            "CreatedAt": "2022-08-31T17:05:13.022+09:00",
            "UpdatedAt": "2022-08-31T21:32:53.394+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 0,
            "bookmark_id": 29,
            "bookmark_url": "https://www.daum.com",
            "bookmark_title": "네이버입니다"
        },
        {
            "CreatedAt": "2022-08-31T17:05:13.418+09:00",
            "UpdatedAt": "2022-08-31T21:32:53.394+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 0,
            "bookmark_id": 30,
            "bookmark_url": "https://www.daum.com",
            "bookmark_title": "네이버입니다"
        },
        {
            "CreatedAt": "2022-08-31T17:05:14.033+09:00",
            "UpdatedAt": "2022-08-31T21:32:53.394+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 0,
            "bookmark_id": 31,
            "bookmark_url": "https://www.daum.com",
            "bookmark_title": "네이버입니다"
        },
        {
            "CreatedAt": "2022-08-31T17:05:11.843+09:00",
            "UpdatedAt": "2022-08-31T21:32:53.394+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 0,
            "bookmark_id": 27,
            "bookmark_url": "https://www.daum.com",
            "bookmark_title": "네이버입니다"
        },
        {
            "CreatedAt": "2022-08-31T17:05:04.29+09:00",
            "UpdatedAt": "2022-08-31T17:05:04.29+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 9,
            "bookmark_id": 26,
            "bookmark_url": "https://www.daum.com",
            "bookmark_title": "네이버입니다"
        }
    ],
    "total_count": 15,
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```

### 2.  POST /bookmarks
본인 소유의 북마크를 생성한다.
#### Request Body
JSON 형식
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| folder_id   | Integer | N(default: 0)   | 해당 북마크를 포함할 본인 소유의 폴더 ID |
| bookmark_url   | String | Y   | 북마크 URL |
| bookmark_title   | String | Y   | 북마크 제목 |
※folder_id는 자신이 소유한 폴더가 아닌 경우 에러 발생
#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 북마크 생성 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E302   | 본인 소유의 폴더가 아님  | 
| E313   | 누락된 Request Body가 존재함 | 

#### 요청 및 응답 예시
[요청]
```
POST http://localhost:8888/bookmarks
```

```json
{
    "folder_id":9,
    "bookmark_url":"https://www.daum.com",
    "bookmark_title":"다음입니다"
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```
### 3.  DELETE /bookmarks
본인 소유의 북마크를 삭제한다.
#### Request Body
JSON 형식
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| bookmark_id   | Integer | Y | 본인 소유의 삭제 대상 북마크 ID |
#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 북마크 삭제 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E302   | 본인 소유의 북마크가 아님  | 

#### 요청 및 응답 예시
[요청]
```
DELETE http://localhost:8888/bookmarks
```

```json
{
    "bookmark_id":32
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```
### 4.  PUT /bookmarks
본인 소유의 북마크 정보를 갱신한다.
#### Request Body
JSON 형식
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| bookmark_id   | Integer | Y | 본인 소유의 수정 대상 북마크 ID |
| folder_id   | Integer | N | 본인 소유의 폴더 ID |
| bookmark_url   | string | N | 북마크 URL |
| bookmark_title   | string | N | 북마크 제목 |
※folder_id는 0으로 변경은 불가 (폴더가 삭제될 때만 0으로 변경됨)
#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 북마크 갱신 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E302   | 본인 소유의 북마크 또는 폴더가 아님  | 
| E303   | 서버 오류 / 요청 데이터의 형식 이상(추후 세분화 예정)  | 


#### 요청 및 응답 예시
[요청]
```
PUT http://localhost:8888/bookmarks
```

```json
{
    "bookmark_id":31,
    "folder_id":7,
    "bookmark_url":"http://123.com",
    "bookmark_title":"title_1"
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```

## 폴더 API

### API 목록

| No   | Method | URI    | 설명  | 권한  | 
| ---- | ------ | ------ | ----- | ----- |
| 1    | GET   | /folders | 본인 폴더 조회 | 사용자 |
| 2    | POST   | /folders | 본인 폴더 생성 | 사용자 |
| 3    | DELETE   | /folders | 본인 폴더 삭제 | 사용자 |
| 4    | PUT   | /folders | 본인 폴더 갱신 | 사용자 |

### 1.  GET /folders
자신의 폴더 리스트를 조건에 맞게 불러온다.
#### Request Query
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| start   | Integer | N(default: 0)   | 조회 시작 순번 |
| count   | Integer | Y(0보다 큰 수)   | 조회 개수 |
| keyword   | String | N  | 폴더 제목으로 리스트 조회 |

#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 폴더 리스트 불러오기 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E312   | 잘못된 형식의 Request Query | 


#### 요청 및 응답 예시
[요청]
```
GET http://localhost:8888/folders?start=0&count=2&keyword=뱁새
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "list": [
        {
            "CreatedAt": "2022-08-31T23:53:23.398+09:00",
            "UpdatedAt": "2022-08-31T23:53:23.398+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 19,
            "folder_title": "뱁새",
            "folder_color": "#123456"
        },
        {
            "CreatedAt": "2022-08-31T23:53:22.913+09:00",
            "UpdatedAt": "2022-08-31T23:53:22.913+09:00",
            "user_email": "m8s8t59rnv@privaterelay.appleid.com",
            "folder_id": 18,
            "folder_title": "뱁새",
            "folder_color": "#123456"
        }
    ],
    "total_count": 10,
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```

### 2.  POST /folders
본인 소유의 폴더를 생성한다.
#### Request Body
JSON 형식
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| folder_title   | String | Y   | 폴더 제목 |
| folder_color   | String | Y   | 폴더 색상 |
#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 폴더 생성 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E313   | 누락된 Request Body가 존재함 | 

#### 요청 및 응답 예시
[요청]
```
POST http://localhost:8888/folders
```

```json
{
    "folder_title":"뱁새",
    "folder_color":"#123456"
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```
### 3.  DELETE /folders
본인 소유의 폴더를 삭제하고, 해당 폴더의 북마크를 기본 소속(folder_id = 0)으로 변경한다.
#### Request Body
JSON 형식
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| folder_id   | Integer | Y(0보다 큰 수) | 본인 소유의 삭제 대상 폴더 ID |
#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 폴더 삭제 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E302   | 본인 소유의 폴더가 아님  | 

#### 요청 및 응답 예시
[요청]
```
DELETE http://localhost:8888/folders
```

```json
{
    "folder_id":7
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```
### 4.  PUT /folders
본인 소유의 폴더 정보를 갱신한다.
#### Request Body
JSON 형식
| 항목   | 타입 | 필수    | 설명 |  
| ---- | ------ | ------ | ----- |
| folder_id   | Integer | Y | 본인 소유의 수정 대상 폴더 ID |
| folder_title   | string | N | 폴더 제목 |
| folder_color   | string | N | 폴더 색상 |
#### Cookie
| 항목        |  설명                 |
| ----------- |  -------------------- |
| access-token | Server로부터 발급받은 만료되지 않은 토큰 |
#### Response Status Code
| Code   | 설명  | 
| ------ | ----- | 
| N200   | 폴더 갱신 완료 | 
| E300   | 비정상적인 Access Token | 
| E301   | 만료된 Access Token  | 
| E302   | 본인 소유의 폴더가 아님  | 


#### 요청 및 응답 예시
[요청]
```
PUT http://localhost:8888/folders
```

```json
{
    "folder_id":19,
    "folder_title":"따오기",
    "folder_color":"#654321"
}
```
[Cookie]
```
access-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im04czh0NTlybnZAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiU25zIjoiYXBwbGUiLCJleHAiOjE2NjE5NTYzNjl9.5o9ZmzdUmvl-qy8W42j1GbLph-SvNOipdrCA_dBGPIM; Path=/; HttpOnly; Expires=Thu, 01 Sep 2022 12:32:49 GMT;
```
[응답]
```json
{
    "api_response": {
        "result_code": "N200",
        "result_message": "OK."
    }
}
```