package handler

import (
	"github.com/gin-gonic/gin"
)



//Router group
func (h *Handler) ShortnerRoute(router *gin.Engine){
	api := router.Group("/shortner")
	{
		api.GET("/:id", h.GetURLbyId)
		api.POST("", h.CreateURL)
	}
}





// func New(log *slog.Logger, urlSaver *service.URLSave) http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request){
// 		log = log.With(
// 			slog.String("request_id", middleware.GetReqID(r.Context())),
// 		)

// 		var req Request
// 		err := render.DecodeJSON(r.Body, &req)
// 		if errors.Is(err, io.EOF) {
// 			// Такую ошибку встретим, если получили запрос с пустым телом
// 			// Обработаем её отдельно
// 			log.Error("request body is empty")
// 			render.JSON(w, r, resp.Response{
// 				Status: resp.StatusError,
// 				Error:  "empty request",
// 			})
// 			return
// 		}
// 		if err != nil{
// 			log.Error("failed to decode request body", sl.Err(err))
// 			render.JSON(w, r, resp.Response{
// 				Status: resp.StatusError,
// 				Error:  "failed to decode request",
// 			})
// 			return
// 		}
// 		log.Info("request body decoded", slog.Any("req", req))

// // Random Aliase creater
// 		alias := req.Alias
//         if alias == "" {
//             alias = random.NewRandomString(aliasLength)
// 		}

// 		id, err := urlSaver.SaveURL(req.URL, alias)
//         if errors.Is(err, storage.ErrURLExists) {
//             // Check created URL
//             log.Info("url already exists", slog.String("url", req.URL))
//             render.JSON(w, r, resp.Error("url already exists"))
//             return
//         }
//         if err != nil {
//             log.Error("failed to add url", sl.Err(err))

//             render.JSON(w, r, resp.Error("failed to add url"))

//             return
//         }

//         log.Info("url added", slog.Int64("id", id))

//         responseOK(w, r, alias)
// 	}
// }