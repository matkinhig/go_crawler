package main

import "fmt"

type Result struct {
	Title       string
	Description string
	Content     string
	Urls        []string
}

func main() {
	fmt.Println("Start golang")

	ch := make(chan string)

	go func() {
		ch <- "https://vnexpress.net/kinh-doanh/chinh-phu-du-dinh-vay-them-gan-500-000-ty-dong-3999848.html"
	}()

	resultChannel := startWorker(ch)

	for result := range resultChannel {
		fmt.Println("find on result : ", result)
		go func() {
			for _, url := range result.Urls {
				ch <- url
			}
		}()
	}
}

func startWorker(ch chan string) chan Result {
	out := make(chan Result)
	go func() {
		for url := range ch {
			go func() {
				result := parseContent(url)
				out <- result
			}()
		}
	}()
	return out
}

func parseContent(url string) Result {
	return Result{
		Title:       "Viet nam",
		Description: "co nhieu tham o",
		Content:     "Bỏ lệ phí môn bài khỏi Luật phí và lệ phí để thúc đẩy việc thành lập doanh nghiệp mới, theo VCCI. Đề xuất trên được Phòng Thương mại và Công nghiệp Việt Nam (VCCI) đưa ra trong góp ý dự thảo Nghị định sửa đổi, bổ sung một số điều của Nghị định số 139/2016 về lệ phí môn bài. Đại diện cơ quan này cho rằng, lệ phí môn bài là khoản tiền mà tổ chức, cá nhân có hoạt động sản xuất, kinh doanh phải nộp hàng năm. Như vậy, cứ có hoạt động kinh doanh là phải đóng, bên cạnh nghĩa vụ với các loại thuế khác phát sinh trong quá trình kinh doanh.",
		Urls:        []string{"https://vnexpress.net/kinh-doanh/chinh-phu-don-luc-lam-san-bay-long-thanh-cao-toc-bac-nam-3999980.html?vn_source=Detail&vn_campaign=Box-XemNhieuNhat&vn_medium=Item-1&vn_term=Desktop&vn_thumb=1", "https://vnexpress.net/kinh-doanh/doanh-nghiep-nha-nuoc-dau-tu-ra-quoc-te-ngay-cang-lo-3998929.html", "https://vnexpress.net/kinh-doanh/vcci-de-nghi-bo-phi-mon-bai-ra-khoi-luat-3998641.html"},
	}

}
