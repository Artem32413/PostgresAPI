package constLog

const (
	ErrDB             = "Ошибка подключения к базе данных "
    ErrDBClose        = "Ошибка закрытия подключения к базе данных "
	ErrDBConnect      = "Ошибка соединения с базой данных "
	ErrDBPing         = "Неудачный пинг "
	ErrNotAllowed     = "Недоступно "
    ErrNotConnect     = "Ошибка отображения данных "
	ErrNotFound       = "Запись не найдена "
	ErrInvalidData    = "Неверные данные "
	ErrInvalidRequest = "Неверный запрос "
	ErrInternal       = "Внутренняя ошибка сервера "
	ErrBadRequest     = "Неверный запрос "
	ErrAuth           = "Ошибка авторизации "
	ErrUnprocessable  = "Неверный формат запроса "
	ErrForbidden      = "Доступ запрещен "
)
