package usecase

import (
	"net/http"

	"TverClip/model"
	"TverClip/repository"

	"golang.org/x/net/html"
)

type ITaskUsecase interface {
	GetAllTasks() ([]model.TaskResponse, error)
	GetTasksByUserID(userID uint) ([]model.TaskResponse, error)
	GetTaskByID(taskID uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(taskID uint, userID uint, newTask model.Task) (model.TaskResponse, error)
	DeleteTask(taskID uint, userID uint) error
}

type TaskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &TaskUsecase{tr}
}

func (tu *TaskUsecase) GetAllTasks() ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks); err != nil {
		return []model.TaskResponse{}, err
	}

	resTasks := []model.TaskResponse{}
	for _, task := range tasks {
		resTasks = append(resTasks, model.TaskResponse{
			ID:             task.ID,
			UserID:         task.UserID,
			Title:          task.Title,
			Description:    task.Description,
			URL:            task.URL,
			IsDone:         task.IsDone,
			OTGImageURL:    task.OTGImageURL,
			OTGDescription: task.OTGDescription,
		})
	}

	return resTasks, nil
}

func (tu *TaskUsecase) GetTasksByUserID(userID uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetTasksByUserID(&tasks, userID); err != nil {
		return []model.TaskResponse{}, err
	}

	resTasks := []model.TaskResponse{}
	for _, task := range tasks {
		resTasks = append(resTasks, model.TaskResponse{
			ID:             task.ID,
			UserID:         task.UserID,
			Title:          task.Title,
			Description:    task.Description,
			URL:            task.URL,
			IsDone:         task.IsDone,
			OTGImageURL:    task.OTGImageURL,
			OTGDescription: task.OTGDescription,
		})
	}

	return resTasks, nil
}

func (tu *TaskUsecase) GetTaskByID(taskID uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskByID(&task, taskID); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:             task.ID,
		UserID:         task.UserID,
		Title:          task.Title,
		Description:    task.Description,
		URL:            task.URL,
		IsDone:         task.IsDone,
		OTGImageURL:    task.OTGImageURL,
		OTGDescription: task.OTGDescription,
	}

	return resTask, nil
}

func (tu *TaskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	ogp, err := extractOGPInfo(task.URL)
	if err != nil {
		return model.TaskResponse{}, err
	}

	task.OTGImageURL = ogp.ImageURL
	task.OTGDescription = ogp.Description

	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:             task.ID,
		UserID:         task.UserID,
		Title:          task.Title,
		Description:    task.Description,
		URL:            task.URL,
		IsDone:         task.IsDone,
		OTGImageURL:    task.OTGImageURL,
		OTGDescription: task.OTGDescription,
	}

	return resTask, nil
}

func (tu *TaskUsecase) UpdateTask(taskID uint, userID uint, newTask model.Task) (model.TaskResponse, error) {
	oldTask := model.Task{}
	if err := tu.tr.GetTaskByID(&oldTask, taskID); err != nil {
		return model.TaskResponse{}, err
	}
	newTask.OTGImageURL = oldTask.OTGImageURL
	newTask.OTGDescription = oldTask.OTGDescription

	if err := tu.tr.UpdateTask(taskID, userID, &newTask); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:             taskID,
		UserID:         userID,
		Title:          newTask.Title,
		Description:    newTask.Description,
		URL:            newTask.URL,
		IsDone:         newTask.IsDone,
		OTGImageURL:    newTask.OTGImageURL,
		OTGDescription: newTask.OTGDescription,
	}

	return resTask, nil
}

func (tu *TaskUsecase) DeleteTask(taskID uint, userID uint) error {
	if err := tu.tr.DeleteTask(taskID, userID); err != nil {
		return err
	}

	return nil
}

type OGPInfo struct {
	ImageURL    string
	Description string
}

// OGPタグをHTMLから抽出する関数
func extractOGPInfo(url string) (*OGPInfo, error) {
	// http.Requestオブジェクトを作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// User-Agentを設定
	req.Header.Set("User-Agent", "facebookexternalhit")

	// HTTPクライアントを作成
	client := &http.Client{}

	// HTTPリクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// HTMLドキュメントをパース
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// OGP情報を格納する変数を初期化
	ogp := OGPInfo{}

	// HTMLノードを再帰的に探索
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			var property, content string
			for _, a := range n.Attr {
				if a.Key == "property" {
					property = a.Val
				} else if a.Key == "content" {
					content = a.Val
				}
			}
			switch property {
			case "og:image":
				ogp.ImageURL = content
			case "og:description":
				ogp.Description = content
			}
			// プロパティと内容を出力
			// if property != "" && content != "" {
			// 	fmt.Printf("Property: %s, Content: %s\n", property, content)
			// }
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return &ogp, nil
}
