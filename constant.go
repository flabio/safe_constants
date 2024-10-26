package safeutils

const ID string = "id"
const MESSAGE string = "message"
const STATUS string = "status"
const DATA string = "data"
const USER string = "user"
const CREATED string = "was successfully created"
const UPDATED string = "was updated successfully"
const REMOVED string = "was successfully removed"
const ERROR_QUERY string = "error query, please try again later"
const ERROR_CREATE string = "error creating"
const ERROR_PARSING_BODY string = "error parsing body"
const ERROR_UPDATE string = "error updating"
const ERROR_DELETE string = "error deleting"
const ERROR_REQUIRED_FIELDS string = " is required."
const EMPTY string = ""
const USER_NOT_FOUND string = "username o password incorrect."
const RECOVER_PANIC string = "Se recuperó de un panic:"
const ID_NO_EXIST string = "The id not exists"
const NAME_ALREADY_EXIST string = "Name already exists"
const EMAIL_ALREADY_EXIST string = "Email already exists"
const ROL_NOT_FOUND string = "Rol not found"
const STATE_NOT_FOUND string = "State not found"
const FIRST_NAME string = "first_name"
const FIRST_SUR_NAME string = "first_sur_name"
const SECOND_SUR_NAME string = "second_sur_name"
const ADDRESS string = "address"
const PHONE string = "phone"
const ZIP_CODE string = "zip_code"
const PROVIDER_NUMBER string = "provider_number"
const STATE_ID string = "state_id"
const ROL_ID string = "rol_id"
const EMAIL string = "email"
const PASSWORD string = "password"
const NAME string = "name"
const TITLE string = "title"
const TIME_HOURS string = "TimeHours"
const COURSE_ID string = "course_id"
const ACTIVE string = "active"
const CITY_ID string = "city_id"
const SCHOOL_ID = "school_id"
const SCHOOL_ID_IS_REQUIRED = "The id school is required."
const COURSE_ID_IS_FIELD_REQUIRED string = "The field course_id is required."
const TITLE_FIELD_IS_REQUIRED string = "The field title is required."
const TIME_HOURS_IS_FIELD_REQUIRED string = "The field time_hours is required."
const SCHOOL_ID_IS_FIELD_REQUIRED = "The field school_id is required."
const FIRST_NAME_IS_FIELD_REQUIRED string = "The field first name is required."
const FIRST_SUR_NAME_IS_FIELD_REQUIRED string = "The field first sur name is required."
const SECOND_SUR_NAME_IS_FIELD_REQUIRED string = "The field second sur name is required."
const ADDRESS_IS_FIELD_REQUIRED string = "The field address is required."
const PHONE_IS_FIELD_REQUIRED string = "The field phone is required."
const ZIP_CODE_IS_FIELD_REQUIRED string = "The field zip code is required."
const PROVIDER_NUMBER_IS_FIELD_REQUIRED string = "The field provider number is required."
const STATE_ID_IS_FIELD_REQUIRED string = "The field state id is required."
const ROL_ID_IS_FIELD_REQUIRED string = "The field rol id is required."
const EMAIL_IS_FIELD_REQUIRED string = "The field email is required."
const PASSWORD_IS_FIELD_REQUIRED string = "The field password is required."
const NAME_FIELD_IS_REQUIRED string = "The name field is required"
const ACTIVE_FIELD_IS_REQUIRED string = "The active field is required"
const CITY_ID_FIELD_IS_REQUIRED string = "The city id field is required"
const FIRST_NAME_IS_MIN_LENGTH_REQUIRED string = "The field first name must have a minimum length of 3 characters."
const FIRST_SUR_NAME_IS_MIN_LENGTH_REQUIRED string = "The field first sur name must have a minimum length of 3 characters."
const SECOND_SUR_NAME_IS_MIN_LENGTH_REQUIRED string = "The field second sur name must have a minimum length of 3 characters."
const ADDRESS_IS_MIN_LENGTH_REQUIRED string = "The field address must have a minimum length of 10 characters."
const PHONE_IS_MIN_LENGTH_REQUIRED string = "The field phone must have a minimum length of 10 characters."
const ZIP_CODE_IS_MIN_LENGTH_REQUIRED string = "The field zip code must have a minimum length of 5 characters."
const STATE_ID_IS_MIN_LENGTH_REQUIRED string = "The field state id "
const FIRST_NAME_IS_REQUIRED string = "The first name is required."
const FIRST_SUR_NAME_IS_REQUIRED string = "The first sur name is required."
const SECOND_SUR_NAME_IS_REQUIRED string = "The second sur name is required."
const ADDRESS_IS_REQUIRED string = "The address is required."
const PHONE_IS_REQUIRED string = "The phone is required."
const ZIP_CODE_IS_REQUIRED string = "The zip code is required."
const PROVIDER_NUMBER_IS_REQUIRED string = "The provider number is required."
const STATE_ID_IS_REQUIRED string = "The state id is required."
const ROL_ID_IS_REQUIRED string = "The rol id is required."
const EMAIL_IS_REQUIRED string = "The email is required."
const PASSWORD_IS_REQUIRED string = "The password is required."
const EMAIL_IS_INVALID string = "The email is invalid."
const CITY_ID_IS_REQUIRED string = "The city id is required."
const NAME_IS_REQUIRED string = "The name is required."
const TITLE_IS_REQUIRED string = "The title is required."
const COURSE_ID_IS_REQUIRED string = "The course id is required."
const TIME_HOURS_IS_REQUIRED string = "The time hours is required."
const PASSWORD_IS_INVALID_LENGTH string = "The password must have a minimum length of 8 characters."
const MSVC_ROL_URL string = "http://3.90.42.175:3001/api/rol"
const MSVC_PARENTESCO_URL string = "http://3.85.174.77:3002/api/parentesco"
const MSVC_EMERGENCY_CONTACT_URL string = "http://3.81.78.83:3003/api/emergency"
const MSVC_CITY_URL string = "http://54.224.197.15:3014/api/cities"
const MSVC_STATES_URL string = "http://54.224.197.15:3014/api/states"
const MSVC_STATES_BY_CITY_URL string = "http://54.224.197.15:3014/api/states/city"
const MSVC_USER_URL string = "http://34.227.88.155:3005/api/user"
const MSVC_AUTH_URL string = "http://34.227.88.155:3005/api/auth"
const MSVC_SCHOOL_URL string = "http://54.91.68.197:3006/api/school"
const MSVC_COURSE_URL string = "http://52.205.201.42:3007/api/course"

const GET string = "GET"
const POST string = "POST"
const PUT string = "PUT"
const DELETE string = "DELETE"
const AUTHORIZATION string = "Authorization"
const BEARER string = "Bearer "
const TOKEN_INVALID string = "Token not provided or invalid"

const LiMIT int = 5
const PAGE string = "page"
const TOTAL_COUNT string = "totalCount"
const PAGE_COUNT string = "pageCount"
const BEGIN string = "begin"

const FILE string = "file"

const FAILED_CREATE string = "Failed to create request"
const SERVICE_NOT_AVAILALE string = "Service not available"
const FIELD_FORM string = "Error al escribir campo de formulario"
const FILE_ERROR_OPEN string = "Error al abrir el archivo"
const FILE_ERROR_CREATE string = "Error al crear archivo en el multipart"
const FILE_COPY_ERROR string = "Error al copiar el archivo"
const FILE_WRITER_ERROR string = "Error al cerrar el multipart writer"
const FILE_ERROR_REQUEST string = "Error al crear la solicitud"
const FILE_ERROR_SERVICE string = "El servicio no está disponible"
const FILE_ERROR_REQUEST_SERVICE string = "Error al leer la respuesta del servicio"
const ENV_ERROR string = "Error loading .env file"

const AWS_BUCKET_NAME string = "bucket-all-is-safe-school"
const AWS_BUCKET_NAME_USER string = "bucket-all-is-safe-user"
const AWS_REGION string = "us-east-1"
const AWS_URL_S3 string = "https://%s.s3.amazonaws.com/%s"
const UPLOADS string = "./uploads/"
const UPLOADS_FILE string = "./uploads/%s"

const DB_DIFF_ID = "id <>?"
const DB_EQUAL_ID = "id=?"
const DB_EQUAL_NAME = "name =?"
const DB_ORDER_DESC = "id desc"
const DB_EQUAL_EMAIL_ID = "email=?"

//pagianation

const LIMIT int = 5
const OFFSET = "offset"
const NAME_IS_FIELD_REQUIRED = "The field name is required."
const MSVC_LANGUEGE_URL = "http://52.205.201.42:3007/api/language"
const MSVC_TOPIC_URL = "http://52.205.201.42:3007/api/topic/"
const DB_EQUAL_USER_ID string = "user_id =?"
const DB_NAME string = "name =?"
const DB_TABLE_NAME string = "Parentesco"
const NAME_FIEDL_IS_REQUIRED = "The field name is required."
const LAST_NAME string = "last_name"
const USER_ID string = "user_id"
const PARENTESCO_ID string = "parentesco_id"
const FIRST_NAME_FIELD_IS_REQUIRED string = "The first name field is required"
const LAST_NAME_FIELD_IS_REQUIRED string = "The last name field is required"
const PHONE_FIELD_IS_REQUIRED string = "The phone field is required"
const USER_ID_FIELD_IS_REQUIRED string = "The user id field is required"
const PARENTESCO_ID_FIELD_IS_REQUIRED string = "The parentesco id field is required"
const LAST_NAME_IS_REQUIRED string = "The last name is required"
const USER_ID_IS_REQUIRED string = "The user ID is required"
const PARENTESCO_ID_IS_REQUIRED string = "The parentest is required"
const ACTIVE_IS_REQUIRED string = "The active is required"
