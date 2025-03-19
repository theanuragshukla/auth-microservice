# Auth-ms Documentation

This document provides an overview of the API endpoints, their descriptions, request/response models, and other relevant details. It also includes information about authentication requirements for certain endpoints.

---

## Table of Contents
1. [Home](#home)
2. [Login](#login)
3. [Profile](#profile)
4. [Signup](#signup)
5. [Token](#token)
6. [Verify](#verify)
7. [Definitions](#definitions)
8. [Responses](#responses)

---

## Home

### **GET** `/`
Returns a welcome message with status code 200.

#### **Operation ID**: `Home`

#### **Tags**:
- home

#### **Responses**:
| Status Code | Description | Schema |
|-------------|-------------|--------|
| 200 | HomeResponse is the response model for the home handler | [`HomeResponse`](#homeresponse) |

---

## Login

### **POST** `/login`
Returns `Tokens` and `uid` if the credentials provided are correct.

#### **Operation ID**: `login`

#### **Parameters**:
| Name | In | Description | Schema |
|------|----|-------------|--------|
| Body | body | Request body containing email and password | - |

#### **Responses**:
| Status Code | Description | Schema |
|-------------|-------------|--------|
| 200 | - | [`LoginResponse`](#loginresponse) |

---

## Profile

### **GET** `/profile`
Returns the user's profile.

#### **Operation ID**: `Profile`

#### **Tags**:
- profile

#### **Authentication**:
- Requires `x-access-token` in the header for authentication.

#### **Responses**:
| Status Code | Description | Schema |
|-------------|-------------|--------|
| 200 | ProfileResponse | [`ProfileResponse`](#profileresponse) |

---

## Signup

### **POST** `/signup`
Returns the user's profile.

#### **Operation ID**: `Signup`

#### **Tags**:
- signup

#### **Responses**:
| Status Code | Description | Schema |
|-------------|-------------|--------|
| 200 | - | [`SignupResponse`](#signupresponse) |

---

## Token

### **GET** `/token`
Refreshes the access token using the refresh token.

#### **Operation ID**: `Token`

#### **Tags**:
- token

#### **Authentication**:
- Requires `x-refresh-token` in the header and `uid` as a query parameter.

#### **Parameters**:
| Name | In | Description | Schema |
|------|----|-------------|--------|
| uid | query | User ID associated with the refresh token | string |
| x-refresh-token | header | Refresh token used to generate a new access token | string |

#### **Responses**:
| Status Code | Description | Schema |
|-------------|-------------|--------|
| 200 | TokenResponse is the response model for the token endpoint | [`TokenResponse`](#tokenresponse) |

---

## Verify

### **GET** `/verify`
Verifies the validity of the access token.

#### **Operation ID**: `Verify`

#### **Tags**:
- verify

#### **Authentication**:
- Requires `x-access-token` in the header for authentication.

#### **Responses**:
| Status Code | Description | Schema |
|-------------|-------------|--------|
| 200 | VerifyResponse is the response model for the verify endpoint | [`VerifyResponse`](#verifyresponse) |

---

## Definitions

### **ProfileResponse**
- **Description**: ProfileResponse is the response model for the profile handler.
- **Go Package**: `auth-ms/handlers`

### **SignupRequest**
- **Go Package**: `auth-ms/handlers`

### **Tokens**
- **Go Package**: `auth-ms/data`

### **User**
- **Description**: User represents the user schema for the database.
- **Go Package**: `auth-ms/data`

---

## Responses

### **HomeResponse**
- **Description**: HomeResponse is the response model for the home handler.
- **Headers**:
  - `msg` (string)
  - `status` (integer, int64)

### **LoginResponse**
- **Description**: -
- **Headers**:
  - `data` (object)
  - `errors` (object)
  - `msg` (string)
  - `status` (boolean)

### **SignupResponse**
- **Description**: -
- **Headers**:
  - `data` (object)
  - `errors` (object)
  - `msg` (string)
  - `status` (boolean)

### **TokenResponse**
- **Description**: TokenResponse is the response model for the token endpoint.
- **Headers**:
  - `data` (object)
  - `msg` (string)
  - `status` (boolean)

### **VerifyResponse**
- **Description**: VerifyResponse is the response model for the verify endpoint.
- **Headers**:
  - `msg` (string)
  - `status` (boolean)

---

### Authentication Details

#### **Authenticated Routes**
- The following routes require authentication via the `x-access-token` header:
  - `/profile`
  - `/verify`

#### **Token Endpoint**
- The `/token` endpoint requires:
  - `x-refresh-token` in the header.
  - `uid` as a query parameter.

---
