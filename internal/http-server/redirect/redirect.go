package redirect
import (
	// "golang.org/x/exp/slog"
)



type URLGetter interface {
    GetURL(alias string) (string, error)
}


// func NewRedirect(urlGetter URLGetter) http.HandlerFunc {
//         const op = "handlers.url.redirect.New"
		
//         // Находим URL по алиасу в БД
//         resURL, err := urlGetter.GetURL(alias)
//         if errors.Is(err, storage.ErrURLNotFound) {
//             // Не нашли URL, сообщаем об этом клиенту
//             log.Info("url not found", "alias", alias)

//             render.JSON(w, r, resp.Error("not found"))

//             return
//         }
//         if err != nil {
//             // Не удалось осуществить поиск
//             log.Error("failed to get url", sl.Err(err))

//             render.JSON(w, r, resp.Error("internal error"))

//             return
//         }

//         log.Info("got url", slog.String("url", resURL))

//         // Делаем редирект на найденный URL
//         http.Redirect(w, r, resURL, http.StatusFound)
//     }
// }