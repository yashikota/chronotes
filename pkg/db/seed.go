package db

import (
	"log"
	"time"

	"gorm.io/gorm"

	model "github.com/yashikota/chronotes/model/v1/db"
)

func Seed(db *gorm.DB) {
	// 既にデータがあるか確認して、存在しなければ挿入
	// if err := db.First(&model.Note{}).Error; err == gorm.ErrRecordNotFound {
	// 	users := []model.User{
	// 		{Name: "John Doe", Email: "john@example.com"},
	// 		{Name: "Jane Smith", Email: "jane@example.com"},
	// 	}
	// 	if err := db.Create(&users).Error; err != nil {
	// 		log.Fatalf("Failed to seed users: %v", err)
	// 	}
	// 	log.Println("Seeded users successfully")
	// } else {
	// 	log.Println("Users already exist, skipping seeding")
	// }

	baseDate, _ := time.Parse("2006-01-02", "2021-09-01")
	id := "01J8BQ16DXVJYJDSPGNKTHSCDV" // 本番環境のデモユーザーID
	// id := "01J8BVCYYJ0SYQ2SC4HXHSFH0N" // ローカル環境のデモユーザーID

	notes := []model.Note{
		{Title: "プロジェクト設計の開始", Content: "新しいプロジェクトの設計を始めるにあたり、まずはユーザーのニーズを徹底的に分析した。その結果、UI/UXの改善が重要だと気づき、デザインに力を入れることにした。この過程で、多くのアイデアが浮かび上がり、実現可能なものをピックアップしていった。", UserID: id, ID: "01J8BQ16DXVJYJDSXFAKXKTHSCDA", Tags: "ui,ux,javascript", CreatedAt: baseDate, UpdatedAt: time.Now()},
		{Title: "会議の振り返り", Content: "昨日の会議で得たフィードバックを元に、プランを修正することにした。具体的には、ユーザーインターフェースの使いやすさをさらに向上させるための改善点が見つかり、チームとのコミュニケーションを重視して進めることが大切だと再認識した。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHSCDA", Tags: "team,agile,devops", CreatedAt: baseDate.AddDate(0, 0, 1), UpdatedAt: time.Now()},
		{Title: "プロジェクト設計の開始", Content: "新しいプロジェクトの設計を始めるにあたり、まずはユーザーのニーズを徹底的に分析した。その結果、UI/UXの改善が重要だと気づき、デザインに力を入れることにした。この過程で、多くのアイデアが浮かび上がり、実現可能なものをピックアップしていった。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS001", Tags: "ui,ux,javascript", CreatedAt: baseDate.AddDate(0, 0, 2), UpdatedAt: time.Now()},
		{Title: "会議の振り返り", Content: "昨日の会議で得たフィードバックを元に、プランを修正することにした。具体的には、ユーザーインターフェースの使いやすさをさらに向上させるための改善点が見つかり、チームとのコミュニケーションを重視して進めることが大切だと再認識した。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS002", Tags: "team,agile,devops", CreatedAt: baseDate.AddDate(0, 0, 3), UpdatedAt: time.Now()},
		{Title: "データベースの最適化", Content: "新しいデータベースのクエリを最適化し、処理速度が劇的に改善されたことに喜びを感じた。特に複雑なデータ構造を扱う際には、効率的なクエリが必要不可欠であると実感し、これからもこのスキルを磨いていきたいと思った。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS003", Tags: "database,sql,optimization", CreatedAt: baseDate.AddDate(0, 0, 4), UpdatedAt: time.Now()},
		{Title: "リモートワークの効率化", Content: "リモートワークでの生産性を上げるために、さまざまなツールを探し始めた。これまで使っていたツールの効果を分析し、どのようにチーム全体の効率を最大化できるかを考えるのはとても有意義な作業だった。今後の成果が楽しみだ。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS004", Tags: "remote,tools,productivity", CreatedAt: baseDate.AddDate(0, 0, 5), UpdatedAt: time.Now()},
		{Title: "映画鑑賞", Content: "友人と映画を観に行った。映画は非常に感動的で、久しぶりにリラックスした時間を過ごすことができた。仕事のストレスを解消するために、こうした時間がどれほど大切かを再確認した。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS005", Tags: "friends,entertainment", CreatedAt: baseDate.AddDate(0, 0, 6), UpdatedAt: time.Now()},
		{Title: "新しい言語の学習", Content: "新しいプログラミング言語を学び始めることにした。初めは戸惑いもあったが、徐々に面白さを感じるようになり、やりがいを感じている。この挑戦が自分のスキルを向上させてくれると信じている。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS006", Tags: "learning,programming,language", CreatedAt: baseDate.AddDate(0, 0, 7), UpdatedAt: time.Now()},
		{Title: "デザインの改善", Content: "デザインのフィードバックを受け取り、改良点を見つけることができた。特に、ユーザーからの意見をもとに改善するプロセスはとても重要で、チーム全体で成長していると感じる。次回のレビューが楽しみだ。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS007", Tags: "design,feedback", CreatedAt: baseDate.AddDate(0, 0, 8), UpdatedAt: time.Now()},
		{Title: "リラックスタイム", Content: "今日は自宅でのんびりと過ごした。好きな本を読んで、心を落ち着けることができた。こうした休息の日があることで、次の仕事に向けての活力を充電できるのだと改めて感じた。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS008", Tags: "relaxation,home", CreatedAt: baseDate.AddDate(0, 0, 9), UpdatedAt: time.Now()},
		{Title: "ライブラリ習得", Content: "新しいライブラリのドキュメントをじっくりと読み込み、使い方をマスターした。特に、実際にプロジェクトに活用できそうな機能が多く、今後の作業が楽しみになった。このライブラリを活用することで、開発がさらにスムーズになるだろう。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS009", Tags: "library,documentation", CreatedAt: baseDate.AddDate(0, 0, 10), UpdatedAt: time.Now()},
		{Title: "アイデア出し", Content: "チームでのブレインストーミングセッションが行われた。新しいアイデアがたくさん出て、創造的なひとときを過ごすことができた。このような会議は、チームの結束力を高めるためにも非常に重要だと感じる。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS010", Tags: "brainstorming,teamwork", CreatedAt: baseDate.AddDate(0, 0, 11), UpdatedAt: time.Now()},
		{Title: "カフェでのひととき", Content: "友人とカフェで過ごし、近況を語り合った。仕事の話や趣味について盛り上がり、とても楽しい時間だった。こうした時間があることで、ストレスを解消し、明日への活力を得ることができる。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS011", Tags: "friends,cafe", CreatedAt: baseDate.AddDate(0, 0, 12), UpdatedAt: time.Now()},
		{Title: "新しいプロジェクトの立ち上げ", Content: "新しいプロジェクトの立ち上げに向けて、チームでのミーティングを行った。プロジェクトの目標やスケジュールを確認し、メンバー全員が理解できるように説明を行った。チーム全体での協力が必要だと感じ、今後の作業に期待を寄せている。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS012", Tags: "project,team,meeting", CreatedAt: baseDate.AddDate(0, 0, 13), UpdatedAt: time.Now()},
		{Title: "python環境の構築", Content: "新しいプロジェクトでpythonを使用することになり、環境構築を行った。特に、ライブラリのインストールやパスの設定が重要だと感じ、丁寧に作業を進めた。この作業を通じて、pythonの基本的な使い方を理解し、今後の開発に活かしていきたいと考えている。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS013", Tags: "python,environment,setup", CreatedAt: baseDate.AddDate(0, 0, 14), UpdatedAt: time.Now()},
		{Title: "pythonの学習,", Content: "環境構築を終えた後、pythonの基本的な文法を学習した。特に、変数や関数の使い方について理解を深め、実際にコードを書いて動作させることで、理論を実践に落とし込むことができた。この学習を通じて、pythonの基礎をしっかりと身につけることができた。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS014", Tags: "python,learning,basics", CreatedAt: baseDate.AddDate(0, 0, 15), UpdatedAt: time.Now()},
		{Title: "pythonを用いてのプログラミング", Content: "基本的な文法を学習した後、実際にpythonを使ってプログラムを作成した。特に、簡単な計算プログラムを作成し、動作確認を行った。この作業を通じて、pythonの使い方を実践的に理解し、今後の開発に活かしていきたいと考えている。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS015", Tags: "python,programming,practice", CreatedAt: baseDate.AddDate(0, 0, 16), UpdatedAt: time.Now()},
		{Title: "未知のエラー", Content: "プログラムを実行した際に、未知のエラーが発生した。特に、エラーメッセージがわかりにくく、原因を特定するのに苦労した。このようなエラーに遭遇することは開発の一部だと理解し、問題解決のスキルを磨いていきたいと考えている。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS016", Tags: "error,debugging,problem-solving", CreatedAt: baseDate.AddDate(0, 0, 17), UpdatedAt: time.Now()},
		{Title: "エラーの解決", Content: "エラーの原因を特定し、解決策を見つけることができた。特に、ログの確認やコードの修正を行い、問題を解決することができた。この作業を通じて、問題解決のスキルを向上させ、今後の開発に活かしていきたいと考えている。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS017", Tags: "error,resolution,debugging", CreatedAt: baseDate.AddDate(0, 0, 18), UpdatedAt: time.Now()},
		{Title: "新しい機能の追加", Content: "プロジェクトに新しい機能を追加することになり、仕様を確認した。特に、ユーザーのニーズに合わせた機能を追加することが重要だと感じ、要件定義を行った。この作業を通じて、ユーザー目線での開発が重要だと再認識した。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS018", Tags: "feature,specification,development", CreatedAt: baseDate.AddDate(0, 0, 19), UpdatedAt: time.Now()},
		{Title: "新しい機能の実装", Content: "仕様を確認した後、新しい機能の実装を行った。特に、デザインや機能の組み合わせにこだわり、ユーザーにとって使いやすい機能を提供することを心がけた。この作業を通じて、ユーザー目線での開発が重要だと再認識した。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS019", Tags: "feature,implementation,development", CreatedAt: baseDate.AddDate(0, 0, 20), UpdatedAt: time.Now()},
		{Title: "新しい機能のテスト", Content: "新しい機能の実装後、テストを行った。特に、ユーザーの操作をシミュレートし、機能が正常に動作するか確認した。この作業を通じて、品質管理の重要性を再認識し、ユーザーにとって価値のある機能を提供することが重要だと感じた。", UserID: id, ID: "01J8BQ16DXVJYJDSPGNKTHS020", Tags: "feature,test,quality", CreatedAt: baseDate.AddDate(0, 0, 21), UpdatedAt: time.Now()},
	}
	if err := db.Create(&notes).Error; err != nil {
		log.Fatalf("Failed to seed posts: %v", err)
	}
	log.Println("Seeded posts successfully")
}
