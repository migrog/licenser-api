# Licenser API

Api Gestor de Licencias (prototipo inicial)

## Ejecutar Localmente (Docker Compose)

Clonar el proyecto

```bash
  git clone https://github.com/migrog/licenser-api.git
```

Ir a la carpeta del proyecto

```bash
  cd licenser-api
```

Ejecutar el comando docker

```bash
  docker-compose up -d
```

## EndPoints

### Seguridad API(localhost:3000)
#### Registrar

```http
  POST /api/register
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | Nombre |
| `email` | `string` | Correo electrónico |
| `password` | `string` | Contraseña |

#### Login

```http
  POST /api/login
```

| Parámetro | Tipo     | Descripción                      |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | Correo electrónico |
| `password`      | `string` | Contraseña |

#### Usuario por Email

```http
  GET /api/users/${email}
```

| Parámetro | Tipo     | Descripción                      |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | Correo electrónico |

### Producto API(localhost:3001)

#### Crear Producto

```http
  POST /api/products
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | Nombre |

#### Obtener Producto

```http
  GET /api/products/${id}
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | Id de producto |


### Licencia API(localhost:3002)

#### Crear Licencia

```http
  POST /api/licenses
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | Correo electrónico de usuario|
| `product_id` | `string` | Id de producto|

#### Verificar Licencia

```http
  GET /api/licenses/verify
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `license_key` | `string` | Clave de licencia|
| `product_id` | `string` | Id de producto|

#### Activar Licencia

```http
  PUT /api/licenses/enable
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `license_key` | `string` | Clave de licencia|
| `product_id` | `string` | Id de producto|

#### Desactivar Licencia

```http
  PUT /api/licenses/disable
```

| Parámetro | Tipo     | Descripción                |
| :-------- | :------- | :------------------------- |
| `license_key` | `string` | Clave de licencia|
| `product_id` | `string` | Id de producto|