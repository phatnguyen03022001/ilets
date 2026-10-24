# SKILL-LEAF-ONTOLOGY — Các năng lực Skill Leaf

> **THẨM QUYỀN PRODUCT CHÍNH THỨC**
>
> Có thẩm quyền trong domain PRODUCT theo `docs/catalog/project.json`. `CONSTITUTION.md` và `OBJECTIVE.md` giữ thẩm quyền riêng; `contracts/**` quản lý machine contract chính xác trong phạm vi của nó. Tham chiếu tới `spec/**` hoặc `design/**` bên dưới, kể cả cách gọi “canonical” trong lịch sử, chỉ nhằm truy nguyên nguồn gốc, không tạo thẩm quyền ngang hàng.

## Mục đích và quyền sở hữu

Xác định **những gì người học phải chứng minh có thể làm được** trong Listening, Reading, Writing và Speaking.

Shard này lưu giữ hệ phân loại năng lực bốn kỹ năng, định danh Skill Leaf, mục tiêu nguyên tử, quan hệ thành viên giữa kỹ năng/thành phần, các phần mở rộng theo biến thể và cạnh điều kiện tiên quyết Skill→Skill. Vocabulary, grammar và phonology là kiến thức hỗ trợ, được giữ trong `docs/product/knowledge-objects.md` với nguồn gốc từ `spec/04-KNOWLEDGE.md`.

Các dạng nhiệm vụ/câu hỏi IELTS chính thức là định danh bài thi bên ngoài, thuộc `spec/02-IELTS-MODEL.md` lịch sử và owner PRODUCT hiện hành. Một Skill Leaf có thể hỗ trợ nhiều dạng chính thức; không được suy ra độ bao phủ dạng câu hỏi/nội dung chỉ từ sự tồn tại của leaf.

Bốn kỹ năng IELTS cấp một chính xác là Listening, Reading, Writing và Speaking.

Danh mục hiện hành có **67 Skill Leaves với ID ổn định**:

- Listening: 11;
- Reading: 12;
- Speaking: 18;
- Writing dùng chung/Academic: 23;
- Phần mở rộng Writing General Training Task 1: 3.

General Training không tạo kỹ năng thứ năm và không nhân đôi năng lực Reading/Writing dùng chung. Leaf theo biến thể chỉ tồn tại khi cấu trúc bài thi bên ngoài khác nhau thực sự.

## Quy tắc Skill Leaf nguyên tử

Skill Leaf thể hiện một năng lực hữu ích độc lập. Chỉ tách tiếp khi node còn nhiều mục tiêu độc lập hoặc không thể dạy, luyện tập, đánh giá và khắc phục một cách độc lập. Dừng khi tách nhỏ hơn không còn cải thiện các đặc tính đó. Không chia năng lực chỉ để phản chiếu mọi nhãn giao diện/dạng câu hỏi của bài thi.

## Trường chính thức của Skill Leaf

Skill Leaf chỉ sở hữu các trường chính thức:

- `id` ổn định;
- tên dễ hiểu;
- kỹ năng IELTS và thành phần;
- biến thể áp dụng nếu không dùng chung;
- một mục tiêu năng lực nguyên tử;
- cạnh tiên quyết Skill→Skill khi một năng lực khác phải có trước.

Những mục sau cố ý **không** thuộc thẩm quyền định danh Skill:

- định danh dạng nhiệm vụ/câu hỏi chính thức → `spec/02-IELTS-MODEL.md` (nguồn lịch sử);
- mức Band liên quan hoặc chất lượng thành thạo Band-N → `spec/05-BANDS.md` (nguồn lịch sử);
- ánh xạ Knowledge Object hỗ trợ → `docs/product/knowledge-objects.md`, kế thừa `spec/04-KNOWLEDGE.md`;
- ánh xạ thực hành → `spec/07-PRACTICE.md` (nguồn lịch sử);
- chiến lược đánh giá/bằng chứng → `spec/08-ASSESSMENT.md` (nguồn lịch sử);
- trạng thái thành thạo của người học → `spec/09-PROGRESSION.md` (nguồn lịch sử);
- vị trí trong chương trình học → `spec/06-CURRICULUM.md` (nguồn lịch sử);
- bao phủ nội dung cụ thể theo dạng/định dạng phụ → `spec/10-CONTENT-MODEL.md` và `design/08-coverage-and-support.md` (nguồn lịch sử);
- lỗi phổ biến, ví dụ, kịch bản khắc phục và chọn scaffold → `spec/10-CONTENT-MODEL.md` (nguồn lịch sử);
- ước lượng mức nhận thức và tải học tập → metadata lập kế hoạch tùy chọn, trừ khi bằng chứng tương lai nâng lên thành chính sách.

## Listening

### Năng lực hiểu nội dung

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `L-COMP-01` | Ý chính / đại ý | Xác định ý trung tâm của lời nói được nghe. |
| `L-COMP-02` | Chi tiết cụ thể | Nắm bắt chính xác những thông tin thực tế được yêu cầu. |
| `L-COMP-03` | Mục đích / thái độ / quan điểm | Suy ra mục đích, thái độ hoặc quan điểm người nói từ ngôn ngữ và cách thể hiện. |
| `L-COMP-04` | Diễn đạt lại và dấu hiệu chuyển ý | Nhận ra cách diễn đạt khác, tín hiệu diễn ngôn và dấu hiệu dẫn đến đáp án. |
| `L-COMP-05` | Xử lý phương án gây nhiễu | Loại bỏ lựa chọn nghe hợp lý nhưng thiếu căn cứ hoặc đã bị thông tin sau thay thế. |
| `L-COMP-06` | Bài nói học thuật dài | Theo dõi độc thoại học thuật dài, dày thông tin mà không mất mạch diễn ngôn. |

### Canonical materializer anchor — Listening main-idea capability

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `L-COMP-01` | Ý chính / đại ý | Xác định ý trung tâm của lời nói được nghe. |

### Chiến lược theo dạng câu hỏi

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `L-QT-01` | Điền form/note/table/flow-chart/summary | Điền chỗ trống có cấu trúc từ audio, tuân thủ bố cục, ngữ pháp và giới hạn từ được nêu. |
| `L-QT-02` | Trắc nghiệm | Chọn đúng một hoặc nhiều phương án đồng thời xử lý phương án gây nhiễu. |
| `L-QT-03` | Ghép nối | Ghép thông tin nghe được với người nói, danh mục hoặc lựa chọn. |
| `L-QT-04` | Gắn nhãn bản đồ/sơ đồ/mặt bằng | Áp dụng ngôn ngữ không gian và chỉ hướng vào hình minh họa. |
| `L-QT-05` | Hoàn thành câu và trả lời ngắn | Tạo câu trả lời ngắn từ audio trong giới hạn đã nêu. |

### Canonical materializer anchor — Listening specific details capability

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `L-COMP-02` | Chi tiết cụ thể | Nắm bắt chính xác những thông tin thực tế được yêu cầu. |

### Canonical materializer anchor — Listening completion question type

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `L-QT-01` | Điền form/note/table/flow-chart/summary | Điền chỗ trống có cấu trúc từ audio, tuân thủ bố cục, ngữ pháp và giới hạn từ được nêu. |

### Các cạnh tiên quyết của Listening

| Leaf phụ thuộc | Skill Leaves tiên quyết |
|---|---|
| `L-COMP-03` | `L-COMP-01` |
| `L-COMP-04` | `L-COMP-01` |
| `L-COMP-05` | `L-COMP-02` |
| `L-COMP-06` | `L-COMP-01`, `L-COMP-04` |
| `L-QT-01` | `L-COMP-02` |
| `L-QT-02` | `L-COMP-01`, `L-COMP-05` |
| `L-QT-03` | `L-COMP-01` |
| `L-QT-04` | `L-COMP-02` |
| `L-QT-05` | `L-COMP-02` |

Phân rã Listening tách năng lực hiểu khỏi chiến lược làm dạng câu hỏi. Việc phân rã phân tích kỹ năng tiếp nhận là suy luận Blueprint dựa trên cấu trúc bài thi chính thức; sự thật về điểm/dạng thi và diễn giải Band vẫn thuộc owner IELTS/Band PRODUCT hiện hành, có nguồn lịch sử ở `spec/02-IELTS-MODEL.md` và `spec/05-BANDS.md`.

## Reading

### Năng lực hiểu nội dung

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `R-COMP-01` | Đọc lướt tìm ý chính | Xác định nhanh ý chính và chủ đề chung. |
| `R-COMP-02` | Đọc quét tìm chi tiết | Nhanh chóng xác định sự kiện hoặc tham chiếu mục tiêu. |
| `R-COMP-03` | Hiểu chi tiết | Hiểu chính xác thông tin tường minh trong bài đọc. |
| `R-COMP-04` | Suy luận / quan điểm tác giả | Suy ra hàm ý, luận điểm, lập trường hoặc ý nghĩa vượt ngoài câu chữ. |
| `R-COMP-05` | Cấu trúc văn bản / mục đích đoạn | Nhận biết bố cục lập luận và chức năng đoạn văn. |
| `R-COMP-06` | Bài đọc dày đặc/trừu tượng/phức tạp | Duy trì mức hiểu trước từ vựng, cú pháp và lập luận khó. |

### Question-type strategies

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `R-QT-01` | Matching headings | Ghép tiêu đề theo ý chính của đoạn, không chỉ trùng từ khóa. |
| `R-QT-02` | True / False / Not Given | Phân biệt thông tin thực tế được chứng minh, bị bác bỏ và không được đề cập. |
| `R-QT-03` | Yes / No / Not Given | Phân biệt đồng tình, không đồng tình và không có thông tin về quan điểm/luận điểm tác giả. |
| `R-QT-04` | Ghép thông tin phân tán | Ánh xạ chính xác thông tin trong matching-information, matching-features và matching-sentence-ending. |
| `R-QT-05` | Điền chỗ trống và trả lời ngắn | Trích xuất, điều chỉnh bằng chứng văn bản theo giới hạn ngữ pháp, cấu trúc và số từ. |
| `R-QT-06` | Trắc nghiệm | Đánh giá lựa chọn theo bằng chứng từ bài đọc, kể cả dạng suy luận. |

### Các cạnh tiên quyết của Reading

| Leaf phụ thuộc | Skill Leaves tiên quyết |
|---|---|
| `R-COMP-03` | `R-COMP-01` |
| `R-COMP-04` | `R-COMP-03` |
| `R-COMP-05` | `R-COMP-01` |
| `R-COMP-06` | `R-COMP-03` |
| `R-QT-01` | `R-COMP-05` |
| `R-QT-02` | `R-COMP-03` |
| `R-QT-03` | `R-COMP-04` |
| `R-QT-04` | `R-COMP-02`, `R-COMP-03` |
| `R-QT-05` | `R-COMP-03` |
| `R-QT-06` | `R-COMP-03`, `R-COMP-04` |

`R-QT-02` và `R-QT-03` vẫn tách biệt vì dạng đầu đánh giá thông tin thực tế, dạng sau đánh giá lập trường tác giả.

`R-QT-04` và `R-QT-05` chủ ý gom các chiến lược nhận thức có thể chuyển giao, thay vì nhân đôi Skill Leaves cho từng dạng giao diện thi chính thức. `spec/02-IELTS-MODEL.md` lưu sự phân biệt các dạng thi trong nguồn lịch sử, còn `spec/10-CONTENT-MODEL.md` lưu yêu cầu kiểm tra độc lập độ bao phủ nội dung chạy được.

Academic và General Training dùng chung các Reading leaves này. Sự khác nhau nằm ở phân bố corpus/ngữ cảnh và chuyển đổi điểm, không phải ontology nhận thức đọc thứ hai. Yêu cầu chuyển giao ngữ cảnh GT thuộc Curriculum hiện hành (nguồn lịch sử `spec/06-CURRICULUM.md`), còn gắn nhãn biến thể/ngữ cảnh cụ thể thuộc Content (nguồn lịch sử `spec/10-CONTENT-MODEL.md`).

## Speaking

### Fluency & Coherence

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `S-FC-01` | Nói liên tục | Tiếp tục nói mà không bị gián đoạn quá mức. |
| `S-FC-02` | Phát triển lượt nói dài | Tạo câu trả lời Part 2 dài, mạch lạc. |
| `S-FC-03` | Dấu hiệu liên kết lời nói | Sử dụng từ nối và dấu hiệu diễn ngôn nói phù hợp. |
| `S-FC-04` | Phát triển chủ đề mạch lạc | Mở rộng, sắp xếp ý hợp lý và đúng trọng tâm. |
| `S-FC-05` | Quản lý ngập ngừng/tự sửa | Không để ngập ngừng hoặc sửa lời cản trở giao tiếp. |

### Lexical Resource

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `S-LR-01` | Từ vựng chủ đề | Duy trì trao đổi bằng từ vựng phù hợp ở các chủ đề thông dụng. |
| `S-LR-02` | Diễn đạt lại khi nói | Nói vòng hoặc diễn đạt khác khi thiếu từ chính xác. |
| `S-LR-03` | Từ ít thông dụng / thành ngữ | Dùng từ ít thông dụng và thành ngữ tự nhiên, phù hợp. |
| `S-LR-04` | Chọn từ chính xác | Chọn từ đúng nghĩa, collocation và sắc thái văn phong. |

### Grammatical Range & Accuracy

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `S-GRA-01` | Độ chính xác câu đơn/ngắn | Tạo cấu trúc nói đơn giản đúng ngay khi chịu tải thời gian thực. |
| `S-GRA-02` | Dạng câu nói phức | Kết hợp hữu ích câu đơn và câu phức khi nói. |
| `S-GRA-03` | Đa dạng và linh hoạt cấu trúc | Dùng linh hoạt nhiều cấu trúc khi nói dài. |
| `S-GRA-04` | Độ chính xác ngữ pháp | Giữ lỗi ngữ pháp đủ hiếm để giao tiếp rõ ràng. |

### Pronunciation

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `S-P-01` | Phát âm vị chính xác | Phát ra âm tiếng Anh đủ rõ để người nghe hiểu. |
| `S-P-02` | Trọng âm từ | Đặt trọng âm từ phù hợp. |
| `S-P-03` | Trọng âm câu và ngữ điệu | Dùng độ nổi bật và biến đổi cao độ để truyền đạt nghĩa. |
| `S-P-04` | Chia cụm / nối âm / nhịp điệu | Chia lời nói thành đơn vị nghĩa tự nhiên, sử dụng nối âm hiệu quả. |
| `S-P-05` | Mức dễ hiểu tổng thể | Người nghe hiểu được mà không phải gắng sức quá mức. |

### Các cạnh tiên quyết của Speaking

| Leaf phụ thuộc | Skill Leaves tiên quyết |
|---|---|
| `S-FC-02` | `S-FC-01` |
| `S-FC-04` | `S-FC-01` |
| `S-FC-05` | `S-FC-01` |
| `S-LR-02` | `S-LR-01` |
| `S-LR-03` | `S-LR-01` |
| `S-GRA-02` | `S-GRA-01` |
| `S-GRA-03` | `S-GRA-02` |
| `S-P-02` | `S-P-01` |
| `S-P-03` | `S-P-02` |
| `S-P-04` | `S-P-01` |
| `S-P-05` | `S-P-01`, `S-P-03` |

Các phụ thuộc Speaking khác là Knowledge Objects hỗ trợ và được phân giải trong `docs/product/knowledge-objects.md`, kế thừa `spec/04-KNOWLEDGE.md`.

## Writing — dùng chung và Academic Task 1

### Task Achievement — Academic Task 1

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `W-TA-01` | Nhận diện đặc điểm hình quan trọng | Chọn xu hướng, khác biệt, giai đoạn, thay đổi hoặc so sánh quan trọng nhất trong hình Academic Task 1. |
| `W-TA-02` | Viết overview Academic | Tóm tắt rõ mẫu/đặc điểm chính, tách khỏi các chi tiết hỗ trợ. |
| `W-TA-03` | Trình bày đặc điểm kèm chi tiết | Bổ trợ đặc điểm đã chọn bằng số liệu chính xác, liên quan hoặc chi tiết không số khi hình yêu cầu. |

### Task Response — phần chung Task 2

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `W-TR-01` | Phân tích đề | Nhận diện mọi yêu cầu đề bài và nhiệm vụ trả lời. |
| `W-TR-02` | Hình thành lập trường rõ ràng | Nêu quan điểm trả lời trực tiếp câu hỏi. |
| `W-TR-03` | Phát triển ý có hỗ trợ | Mở rộng ý chính bằng lý do, ví dụ hoặc bằng chứng liên quan. |
| `W-TR-04` | Duy trì tính liên quan | Bám sát đề, tránh kéo dài hoặc lặp ý. |

### Coherence & Cohesion

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `W-CC-01` | Chia đoạn | Gom nội dung thành đoạn hợp lý với mục đích rõ. |
| `W-CC-02` | Tiến triển logic | Sắp xếp thông tin, ý tưởng theo mạch tổng thể rõ ràng. |
| `W-CC-03` | Phương tiện liên kết | Dùng từ/cấu trúc liên kết chính xác, không máy móc. |
| `W-CC-04` | Quy chiếu và thay thế | Nối ý, giảm lặp bằng quy chiếu/thay thế rõ ràng. |

### Lexical Resource

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `W-LR-01` | Từ vựng theo chủ đề | Dùng đủ từ vựng phù hợp với chủ đề đề bài. |
| `W-LR-02` | Collocation | Dùng các tổ hợp từ tự nhiên, chính xác. |
| `W-LR-03` | Diễn đạt lại | Diễn đạt lại đề, số liệu và ý mà không làm méo nghĩa. |
| `W-LR-04` | Từ chính xác / ít thông dụng / thành ngữ | Dùng vốn từ nâng cao khi tăng độ chính xác và vẫn phù hợp. |
| `W-LR-05` | Chính tả và cấu tạo từ | Kiểm soát chính xác chính tả và dạng từ trong họ. |

### Grammatical Range & Accuracy

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `W-GRA-01` | Độ chính xác câu đơn | Tạo câu đơn đúng cấu trúc. |
| `W-GRA-02` | Câu ghép | Phối hợp chính xác các mệnh đề độc lập. |
| `W-GRA-03` | Câu phức | Sử dụng hiệu quả mệnh đề phụ và cấu trúc phức. |
| `W-GRA-04` | Thì và thể | Chọn thì/thể phù hợp mục đích giao tiếp. |
| `W-GRA-05` | Mạo từ và từ hạn định | Kiểm soát mạo từ, từ hạn định và danh từ đếm được. |
| `W-GRA-06` | Dấu câu | Dùng dấu câu giúp rõ cấu trúc ngữ pháp và diễn ngôn. |
| `W-GRA-07` | Linh hoạt và đa dạng cấu trúc | Dùng nhiều cấu trúc mà không đánh đổi độ chính xác. |

## Writing — phần mở rộng General Training Task 1

General Training Task 1 không dùng các leaf đặc điểm hình Academic. Dạng này có ba leaf riêng và vẫn dùng chung năng lực liên kết, từ vựng, ngữ pháp và diễn đạt.

| ID | Năng lực | Mục tiêu nguyên tử |
|---|---|---|
| `W-GT1-01` | Phân tích tình huống / người nhận / mục đích | Xác định quan hệ với người nhận, mục đích giao tiếp và mọi gạch đầu dòng yêu cầu trong đề. |
| `W-GT1-02` | Kiểm soát văn phong và quan hệ trong thư | Chọn, duy trì văn phong thân mật, bán trang trọng hoặc trang trọng phù hợp người nhận/mục đích, gồm mở/kết thư và giọng điệu. |
| `W-GT1-03` | Đáp ứng mục đích và yêu cầu | Đạt mục đích giao tiếp của thư và bao phủ từng gạch đầu dòng bằng thông tin liên quan, phát triển đủ. |

Các leaf này chỉ hoạt động khi biến thể Writing Task 1 là General Training. Chúng được đánh giá theo tiêu chí Writing thông thường và không tạo bản sao định nghĩa từ vựng/ngữ pháp/liên kết.

### Các cạnh tiên quyết của Writing

| Leaf phụ thuộc | Skill Leaves tiên quyết |
|---|---|
| `W-TA-02` | `W-TA-01` |
| `W-TA-03` | `W-TA-01` |
| `W-TR-02` | `W-TR-01` |
| `W-TR-03` | `W-TR-02` |
| `W-TR-04` | `W-TR-01` |
| `W-CC-02` | `W-CC-01` |
| `W-CC-03` | `W-CC-02` |
| `W-CC-04` | `W-CC-02` |
| `W-LR-02` | `W-LR-01` |
| `W-LR-03` | `W-LR-01` |
| `W-LR-04` | `W-LR-01` |
| `W-GRA-02` | `W-GRA-01` |
| `W-GRA-03` | `W-GRA-01` |
| `W-GRA-07` | `W-GRA-02`, `W-GRA-03` |
| `W-GT1-02` | `W-GT1-01` |
| `W-GT1-03` | `W-GT1-01` |

Kiến thức về văn phong GT được phân giải qua `K-VOC-040` trong `docs/product/knowledge-objects.md`, kế thừa `spec/04-KNOWLEDGE.md`; năng lực viết thư vẫn là Skill vì người học phải chọn và dùng kiến thức đó phù hợp với nhiệm vụ giao tiếp.

Các Writing leaves chỉ có điều kiện Knowledge được phân giải trong `docs/product/knowledge-objects.md`. Độ dài và tuân thủ định dạng thi là ràng buộc nhiệm vụ bên ngoài, không phải Skill Leaves.

## Bất biến về bao gồm theo biến thể

Với Writing Task 1 `Academic`, năng lực Task 1 bắt buộc gồm `W-TA-01`, `W-TA-02` và `W-TA-03`, cộng những Writing leaves dùng chung phù hợp.

Với Writing Task 1 `General Training`, năng lực Task 1 bắt buộc gồm `W-GT1-01`, `W-GT1-02` và `W-GT1-03`, cộng Writing leaves dùng chung; các leaf trực quan `W-TA-*` không phải điều kiện tiên quyết GT Task 1.

Task 2 dùng `W-TR-*` chung và các tiêu chí/năng lực Writing chung cho cả hai biến thể, trừ khi owner IELTS chính thức (nguồn lịch sử `spec/02-IELTS-MODEL.md`) sau này xác định khác biệt cấu trúc thật sự cần leaf riêng.

## Bất biến về quyền sở hữu

- `spec/02-IELTS-MODEL.md` là nguồn lịch sử về định danh dạng nhiệm vụ/câu hỏi IELTS chính thức;
- `docs/product/knowledge-objects.md`, kế thừa `spec/04-KNOWLEDGE.md`, là owner duy nhất của ánh xạ Skill→Knowledge trong biểu diễn PRODUCT kế nhiệm;
- `spec/05-BANDS.md` là nguồn lịch sử về ngưỡng chất lượng/đầu ra Band-N, gồm phần mở rộng Task 1 theo biến thể;
- `spec/06-CURRICULUM.md` là nguồn lịch sử về sắp xếp Skill Leaves nhưng không được định nghĩa lại;
- `spec/07-PRACTICE.md` là nguồn lịch sử về ánh xạ Practice Types sang leaves nhưng không được định nghĩa lại;
- `spec/08-ASSESSMENT.md` là nguồn lịch sử về đo lường leaves nhưng không được định nghĩa lại;
- `spec/09-PROGRESSION.md` là nguồn lịch sử về trạng thái người học theo leaf ID nhưng không được định nghĩa lại năng lực đích;
- `spec/10-CONTENT-MODEL.md` là nguồn lịch sử về bao phủ nội dung chạy được theo dạng/ngữ cảnh/định dạng phụ mà không tạo định nghĩa Skill song song.

Không owner nào ở bước sau được tạo định nghĩa Skill Leaf song song.
