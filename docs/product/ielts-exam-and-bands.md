# IELTS-EXAM-BANDS Ngữ nghĩa kỳ thi IELTS bên ngoài và Band

> **THẨM QUYỀN PRODUCT CHÍNH THỨC**
>
> Có thẩm quyền trong domain PRODUCT theo `docs/catalog/project.json`. `CONSTITUTION.md` và `OBJECTIVE.md` giữ thẩm quyền riêng; `contracts/**` có thẩm quyền đối với machine contract chính xác trong phạm vi của nó. Các tham chiếu tới `spec/**` hoặc `design/**`, kể cả cách gọi “canonical” trong lịch sử, chỉ là nguồn gốc trước chuyển đổi, không tạo thẩm quyền ngang hàng.

## Cấu trúc bài thi IELTS bên ngoài

IELTS chuẩn có hai biến thể: IELTS Academic và IELTS General Training. Listening và Speaking dùng chung; Reading và Writing khác khi bài thi bên ngoài có khác biệt. Hình thức tổ chức thi là chiều độc lập, không phải biến thể IELTS khác.

IELTS báo cáo bốn điểm Band: Listening, Reading, Writing và Speaking. Overall Band là trung bình cộng của bốn điểm này, làm tròn tới Band nguyên hoặc nửa Band gần nhất; `.25` làm tròn lên nửa Band kế tiếp và `.75` lên Band nguyên kế tiếp. Overall không có nghĩa năng lực các kỹ năng đồng đều.

### Các định danh dạng thi chính thức ổn định

Những ID này chỉ các dạng bài thi bên ngoài, không phải Skills, Practice Types, tính năng sản phẩm hay enum trên wire. ID ổn định chừng nào dạng bên ngoài tương ứng chưa thay đổi đáng kể; độ bao phủ được theo dõi riêng với định danh Skill.

| ID | Dạng thi bên ngoài |
|---|---|
| `IELTS-L-QF-01` | Listening — trắc nghiệm |
| `IELTS-L-QF-02` | Listening — ghép nối |
| `IELTS-L-QF-03` | Listening — gắn nhãn mặt bằng / bản đồ / sơ đồ |
| `IELTS-L-QF-04` | Listening — điền form / note / table / flow-chart / summary |
| `IELTS-L-QF-05` | Listening — hoàn thành câu |
| `IELTS-L-QF-06` | Listening — câu hỏi trả lời ngắn |
| `IELTS-R-QF-01` | Reading — trắc nghiệm |
| `IELTS-R-QF-02` | Reading — True / False / Not Given |
| `IELTS-R-QF-03` | Reading — Yes / No / Not Given |
| `IELTS-R-QF-04` | Reading — ghép thông tin |
| `IELTS-R-QF-05` | Reading — ghép tiêu đề |
| `IELTS-R-QF-06` | Reading — ghép đặc điểm |
| `IELTS-R-QF-07` | Reading — ghép vế kết câu |
| `IELTS-R-QF-08` | Reading — hoàn thành câu |
| `IELTS-R-QF-09` | Reading — điền summary / note / table / flow-chart |
| `IELTS-R-QF-10` | Reading — điền nhãn sơ đồ |
| `IELTS-R-QF-11` | Reading — câu hỏi trả lời ngắn |
| `IELTS-W-A-T1` | Academic Writing Task 1 — thông tin trực quan |
| `IELTS-W-GT-T1` | General Training Writing Task 1 — viết thư |
| `IELTS-W-T2` | Writing Task 2 — bài luận; năng lực học cốt lõi dùng chung, đề cụ thể vẫn thuộc phạm vi biến thể |
| `IELTS-S-P1` | Speaking Part 1 — phỏng vấn/chủ đề quen thuộc |
| `IELTS-S-P2` | Speaking Part 2 — lượt nói dài cá nhân |
| `IELTS-S-P3` | Speaking Part 3 — thảo luận mở rộng |

#### Canonical materializer anchor — Listening multiple-choice family

| ID | Dạng thi bên ngoài |
|---|---|
| `IELTS-L-QF-01` | Listening — trắc nghiệm |

#### Canonical materializer anchor — Listening completion family

| ID | Dạng thi bên ngoài |
|---|---|
| `IELTS-L-QF-04` | Listening — điền form / note / table / flow-chart / summary |

#### Canonical materializer anchor — Reading classification families

| ID | Dạng thi chính thức |
|---|---|
| `IELTS-R-QF-02` | Xác định thông tin — True / False / Not Given |
| `IELTS-R-QF-03` | Xác định quan điểm/luận điểm tác giả — Yes / No / Not Given |

#### Canonical materializer anchor — Reading matching headings family

| ID | Dạng thi chính thức |
|---|---|
| `IELTS-R-QF-05` | Reading — ghép tiêu đề |

Một Skill Leaf có thể phục vụ nhiều dạng thi chính thức, và một dạng có thể yêu cầu nhiều Skills. Các định dạng trình bày hình Academic Task 1 là chiều bao phủ nội dung/trình bày, không phải nhiệm vụ Writing chấm điểm bổ sung. Speaking vẫn là cấu trúc Band đánh giá tổng thể dù Parts 1–3 có ID dạng thi ổn định để theo dõi bao phủ.

## Hợp đồng về mức độ chắc chắn của sự thật bên ngoài

Mọi phát biểu mang ý nghĩa IELTS hoặc Band bên ngoài sử dụng lớp mạnh nhất mà bằng chứng nguồn thực sự hỗ trợ:

- **Official Evidence** — sự thật, mô tả tiêu chí, quy tắc/hướng dẫn chấm, yêu cầu định dạng hoặc sự thật bên ngoài khác được nguồn IELTS có thẩm quyền xác lập trực tiếp.
- **Evidence-Based Interpretation** — diễn giải giáo dục gần sát, có cơ sở từ nguồn chính thức hoặc bằng chứng quan trọng khác, nhưng không phải phát biểu trực tiếp của IELTS.
- **Blueprint Inference** — chi tiết sư phạm do blueprint sản phẩm này đưa thêm ngoài những gì IELTS công bố tường minh.

Evidence-Based Interpretation và Blueprint Inference không bao giờ được trình bày như Official Evidence trực tiếp. Official Evidence có thể thay đổi phải giữ nguồn gốc kèm ngày để phân biệt baseline lịch sử đã lưu với thông tin ngoài được kiểm tra lại hiện tại.

### Cập nhật sự thật IELTS qua owner trước

Khi sự thật IELTS bên ngoài có thẩm quyền thay đổi:

1. cập nhật owner sự thật bên ngoài này trước;
2. làm mới bằng chứng/nguồn gốc có ngày cho sự thật thay đổi;
3. rà soát mọi hệ quả downstream liên quan đến Skill, Band, Curriculum, Assessment, Content, coverage, behavior và system design trước khi đổi hành vi suy ra;
4. chỉ giữ ID dạng thi chính thức ổn định nếu dạng bên ngoài không thay đổi đáng kể; và
5. không giữ sự thật bên ngoài đã lỗi thời chỉ vì implementation, content, configuration hoặc artifact downstream khác đã mã hóa nó.

Owner downstream được diễn giải hoặc vận hành sự thật bên ngoài trong phạm vi của mình, nhưng không ghi đè owner này hoặc âm thầm định nghĩa lại năng lực IELTS.

### Baseline về hình thức thi và hành chính

Các sự thật bên ngoài được owner cũ kiểm tra lần cuối với tài liệu IELTS chính thức ngày 2026-08-22. Biểu diễn kế nhiệm giữ baseline đã khóa đó; sự thật bên ngoài đang biến đổi phải được kiểm tra lại trước implementation/release.

- Thi máy tính tại trung tâm là baseline tổ chức IELTS chuẩn hiện tại. Listening, Reading và Writing tương tác qua máy tính; Speaking vẫn là tương tác với giám khảo người thật. Hình thức thi không làm thay đổi cấu trúc năng lực ngôn ngữ hoặc cách diễn giải Band.
- Một số thị trường cung cấp Writing on Paper trong IELTS on computer. Listening/Reading vẫn làm trên máy; đề Writing hiện trên màn hình và bài viết trả lời bằng tay. Dạng bài, thời gian, cách chấm và tiêu chí không đổi. Hiện chưa cung cấp cho IELTS for UKVI, và Writing One Skill Retake phải dùng cùng hình thức Writing với bài thi gốc đủ điều kiện khi tuyến đó áp dụng.
- IELTS Online chỉ dành cho Academic tại các quốc gia hỗ trợ. Listening/Reading/Writing làm từ xa; Speaking là cuộc gọi video trực tiếp bảo mật với giám khảo. Tổ chức nhận kết quả quyết định chấp nhận hay không; tuyến này hiện không dành cho mục đích visa/di trú và không có General Training trên IELTS Online.
- UKVI Academic/GT dùng lại cấu trúc ngôn ngữ tương ứng, thêm ràng buộc hành chính/bảo mật/địa điểm; những ràng buộc này không tạo chương trình học hay định nghĩa Band khác.
- One Skill Retake sử dụng lại một kỹ năng hiện có. Điều kiện đã khóa gồm: bài thi IELTS on computer đầy đủ ban đầu đủ điều kiện ở trung tâm tham gia; được thi lại một kỹ năng một lần cho mỗi bài đầy đủ gốc; thi lại trong 60 ngày; khả năng cung cấp tùy địa điểm; tổ chức nhận có thể quy định điều kiện chấp nhận; hiện chấp nhận UKVI cho bài Academic/GT trên máy tại những địa điểm được chọn; Writing phải khớp hình thức Writing gốc đủ điều kiện khi áp dụng.

### Sự thật về thời gian, cấu trúc và chấm điểm

Toàn bài thi khoảng 2h45m: Listening khoảng 30m, Reading 60m, Writing 60m, Speaking khoảng 11–14m.

Listening có 40 câu trong bốn phần, mỗi phần 10 câu; audio chỉ nghe một lần. Ngữ cảnh tăng dần từ hội thoại đời sống/xã hội, độc thoại đời sống/xã hội, hội thoại giáo dục/đào tạo tới độc thoại học thuật. Mỗi câu đúng một điểm, điểm thô /40 chuyển thành Band. Các mốc trung bình hiện tại: Band 5≈16, 6≈23, 7≈30, 8≈35; số câu có thể khác nhẹ theo phiên bản đề, không phải bảng quy đổi cố định vĩnh viễn.

Reading có 40 câu trong 60m. Academic dùng ba bài đọc định hướng học thuật. Các mốc trung bình Academic hiện tại: Band 5≈15, 6≈23, 7≈30, 8≈35. General Training dùng Section 1 các văn bản đời sống, Section 2 văn bản công việc, Section 3 văn bản phổ thông dài hơn, với hướng dẫn hiện tại về tổng độ dài khoảng 2150–2750 từ; các mốc GT trung bình hiện tại: Band 4≈15, 5≈23, 6≈30, 7≈35. Không được hoán đổi bảng quy đổi Academic và GT.

Writing có hai nhiệm vụ trong 60m. Bốn tiêu chí là Task Achievement (Task 1) / Task Response (Task 2), Coherence and Cohesion, Lexical Resource, Grammatical Range and Accuracy. Các tiêu chí có trọng số bằng nhau trong mỗi Task; Task 2 nặng gấp đôi Task 1. Academic Task 1 trình bày thông tin trực quan, tối thiểu 150 từ, khoảng 20m. GT Task 1 viết thư, tối thiểu 150 từ, khoảng 20m, với ba chức năng nội dung/gạch đầu dòng cho sẵn và phong cách thân mật/bán trang trọng/trang trọng tùy quan hệ. Task 2 viết luận, tối thiểu 250 từ, khoảng 40m.

Speaking là tương tác với giám khảo người thật, được ghi âm, dùng chung Academic/GT. Bốn tiêu chí Fluency and Coherence, Lexical Resource, Grammatical Range and Accuracy, Pronunciation có trọng số bằng nhau. Có thể thi trực tiếp hoặc với IELTS Online Academic qua video trực tiếp bảo mật; cấu trúc năng lực đánh giá vẫn như nhau.

## Ngưỡng học tập Bands 3–9

Bands 3–9 là khoảng học có cấu trúc. Bands 0–2 vẫn là ranh giới chẩn đoán/bên ngoài, không có chương trình học chi tiết chính thức. Các mô tả bổ sung kỹ năng tiếp nhận dưới đây thuộc Blueprint Inference vì điểm Listening/Reading công khai chủ yếu dựa vào điểm thô; ngưỡng kỹ năng tạo sản phẩm là diễn giải sát tiêu chí phân tích đã công bố. Ngưỡng Band xác định chất lượng yêu cầu chứ không phải số lượng bằng chứng; Assessment quyết định tính đủ.

### Band 3 — đầu vào có cấu trúc

- Listening: lấy nghĩa rời rạc/đơn giản từ lời nói hằng ngày dễ đoán; nhận ra từ quen, mệnh đề đơn giản, một số chi tiết cơ bản và bắt đầu phân biệt ý chính với từ đơn. Có thể còn khoảng trống hiểu lớn, theo dõi lời nói dài thiếu tin cậy, nhận ra paraphrase/dấu hiệu chuyển ý yếu và chưa theo kịp độc thoại học thuật.
- Reading: tìm thông tin tường minh rời rạc trong văn bản dễ tiếp cận bằng khớp từ/đọc quét cơ bản và bắt đầu nắm chủ đề chung đơn giản. Suy luận đáng tin cậy, tích hợp ý chính và đọc văn bản dày/phức tạp chưa thuộc ngưỡng mong đợi.
- Writing: thử làm đủ hai nhiệm vụ với độ dài có thể chấm, thử cấu trúc câu, truyền đạt một phần nghĩa liên quan và tránh hoàn toàn chỉ chép thuộc lòng. Có thể thường xuyên sai ngữ pháp/chính tả/bố cục/từ vựng/kiểm soát nhiệm vụ; Academic Task 1 chỉ chọn/tả được một phần đặc điểm, GT Task 1 chỉ kiểm soát một phần người nhận/mục đích/gạch đầu dòng/văn phong.
- Speaking: truyền đạt nghĩa cơ bản về chủ đề quen thuộc/cá nhân, thử câu đơn, nối đáp án tối thiểu và phần nào giao tiếp được dù dừng nhiều/tìm từ. Có thể còn nhiều lỗi ngữ pháp, từ vựng và phát âm.

### Band 4 — kiểm soát chức năng cơ bản

- Listening: nhận ra ý chính của bản ghi đơn giản hơn, nắm chi tiết thông dụng, làm câu điền/trả lời ngắn cơ bản; paraphrase, xử lý gây nhiễu và nghe học thuật dày vẫn chưa ổn định.
- Reading: đọc quét chi tiết, đọc lướt ý chính cơ bản, làm điền/trả lời ngắn đơn giản; suy luận, lập trường tác giả và văn bản trừu tượng dày chưa đạt ngưỡng.
- Writing: đáp ứng tối thiểu mỗi nhiệm vụ, thể hiện lập trường Task 2 có thể nhận ra, có một số câu đơn đúng, chia đoạn và liên kết cơ bản. Academic Task 1 nhận diện vài đặc điểm hình quan trọng; GT Task 1 nhận ra người nhận/mục đích, cố đáp ứng gạch đầu dòng và dùng ít nhiều văn phong phù hợp quan hệ. Lỗi/lặp thường xuyên vẫn chấp nhận nếu người đọc hiểu được nghĩa.
- Speaking: thảo luận cơ bản chủ đề quen, nối câu đơn, có một số phát ngôn ngắn đúng và thử trọng âm/ngữ điệu cơ bản, nhìn chung dễ hiểu; vẫn có thể dừng lâu, lặp, sai thường xuyên và trượt phát âm.

### Band 5 — kiểm soát đang phát triển

- Listening: theo dõi ý chính/chi tiết liên quan trong bối cảnh xã hội quen và một phần giáo dục, nhận biết paraphrase thông dụng, bắt đầu xử lý gây nhiễu và nhiều dạng câu hỏi hơn; chi tiết học thuật dày/paraphrase phức tạp vẫn thiếu ổn định.
- Reading: nhận diện ý chính và thông tin tường minh chi tiết, bắt đầu suy luận/cấu trúc và dùng chiến lược vượt ngoài điền chỗ trống; lập luận trừu tượng/dày và lập trường tinh tế vẫn chưa tin cậy. GT bắt đầu chuyển giao giữa ngữ cảnh đời sống/nơi làm việc.
- Writing: đáp ứng các yêu cầu chính nhưng chưa đầy đủ, nêu lập trường Task 2 ít phát triển, có bố cục nền, thử ngữ pháp phức với độ chính xác hạn chế và vốn từ vừa đủ. Academic Task 1 thử overview và đặc điểm chọn lọc. GT Task 1 xác định mục đích, bao phủ phần lớn gạch đầu dòng và dùng văn phong tương đối hợp lý dù thiếu nhất quán. Lỗi đáng kể, liên kết máy móc và phát triển ý thiếu vẫn tương thích mức này.
- Speaking: duy trì nói dù lặp/tự sửa/chậm, thảo luận chủ đề quen/lạ với khả năng dùng từ linh hoạt còn hạn chế, thử paraphrase, kiểm soát cấu trúc cơ bản dù câu phức dễ sai.

### Band 6 — thể hiện rõ ràng, đủ năng lực

- Listening: theo dõi ý chính/chi tiết cụ thể ở phần lớn bài, xử lý paraphrase/dạng câu hỏi phổ biến, bắt đầu theo kịp lời nói học thuật dài và xử lý nhiều phương án gây nhiễu dù chưa luôn nhất quán.
- Reading: hiểu ý chính và thông tin tường minh chi tiết, xử lý phần lớn dạng câu hỏi chính thức, có một số suy luận tin cậy và đọc văn bản dài/dày hơn dù sắc thái/trừu tượng còn khó. GT chuyển giao tốt qua phần đời sống, nơi làm việc và văn bản phổ thông dài.
- Writing Task 2 đáp ứng các phần chính của đề, trình bày lập trường liên quan và phát triển ý chính đủ. Chất lượng chung có tiến triển tương đối mạch lạc, vốn từ đủ cho nhiệm vụ, kết hợp câu đơn/phức, lỗi ít khi chặn giao tiếp. Academic Task 1 đáp ứng yêu cầu chính, chọn đặc điểm hình liên quan, thử overview phù hợp và hỗ trợ mô tả bằng số liệu/chi tiết. GT Task 1 xác định đúng người nhận/mục đích, giải quyết mọi gạch đầu dòng bằng thông tin liên quan đủ, văn phong/bố cục thư nhìn chung phù hợp và tránh lựa chọn văn phong gây hại đáng kể đến mục đích.
- Speaking: trao đổi chủ đề dài, nói lượt dài dù đôi khi giảm mạch lạc, nhìn chung paraphrase được, dùng câu đơn/phức, có trọng âm/ngữ điệu hiệu quả và nhìn chung người nghe hiểu.

### Band 7 — kiểm soát tốt, linh hoạt

- Listening: theo dõi lời nói dài đáng tin cậy, xử lý nội dung học thuật/trừu tượng, thường xuyên nhận ra paraphrase/dấu hiệu chuyển ý, loại phần lớn gây nhiễu, chỉ bỏ lỡ sắc thái/chi tiết khó thỉnh thoảng.
- Reading: đọc bài dài/phức tạp tin cậy, suy luận nghĩa/lập trường, hiểu tổ chức văn bản và giải bài ghép/lập trường khó một cách ổn định. GT đáng tin cậy qua đủ Section 1/2/3.
- Writing Task 2 nêu lập trường rõ và phát triển đủ, mở rộng/hỗ trợ ý và giữ trọng tâm. Chất lượng chung tiến triển logic rõ, chia đoạn nhìn chung hiệu quả, liên kết/quy chiếu linh hoạt, có từ ít thông dụng/chính xác, câu phức đa dạng và nhiều câu không lỗi. Academic Task 1 có overview rõ, nêu xu hướng/khác biệt/đặc điểm chính và tổ chức chi tiết. GT Task 1 đạt rõ mục đích, phát triển mọi gạch đầu dòng, giữ quan hệ/văn phong phù hợp chỉ đôi khi sai, dùng mở/kết thư và bố cục đúng tình huống giao tiếp.
- Speaking: nói lượt dài không quá gắng sức, phát triển chủ đề mạch lạc, dùng dấu hiệu diễn ngôn linh hoạt, paraphrase hiệu quả, nhiều cấu trúc hữu ích với câu đúng thường xuyên và phát âm hiệu quả dù đôi khi trượt.

### Band 8 — kiểm soát rất mạnh

- Listening: hiểu lời nói chi tiết, phức tạp, trừu tượng với ít khó khăn, xử lý gây nhiễu/paraphrase tự tin, chỉ hiếm sai chi tiết tinh tế/dày đặc.
- Reading: hiểu văn bản chi tiết, phức tạp, trừu tượng ít khó khăn; xử lý suy luận, lập trường, cấu trúc và dạng khó tin cậy, hiếm sai lập luận tinh tế. Khác biệt ngữ cảnh/văn phong GT không còn gây bất ổn đáng kể.
- Writing Task 2 đáp ứng đề phù hợp/đầy đủ, nêu lập trường phát triển tốt và hỗ trợ ý hiệu quả. Chất lượng chung dễ theo dõi, liên kết được kiểm soát tốt, từ vựng rộng/linh hoạt/chính xác, cấu trúc đa dạng, phần lớn câu không lỗi và chỉ có lỗi không hệ thống thỉnh thoảng. Academic Task 1 đáp ứng nhiệm vụ phù hợp/đầy đủ, chọn/nhấn mạnh/minh họa đặc điểm chính khéo léo. GT Task 1 đáp ứng mục đích/nội dung phù hợp/đầy đủ, kiểm soát văn phong linh hoạt/nhất quán, chọn từ/giọng điệu/bố cục/mở-kết thư nhạy cảm người nhận và chỉ đôi khi trượt không hệ thống.
- Speaking: nói trôi chảy với ngập ngừng chủ yếu do nội dung, từ vựng và cấu trúc rộng/linh hoạt/chính xác, hầu hết câu không lỗi, duy trì nhịp điệu/trọng âm/ngữ điệu/độ dễ hiểu hiệu quả.

### Band 9 — thành thạo mức trần

- Listening: hiểu gần như mọi nội dung nói, kể cả tinh tế, trừu tượng, chi tiết và dài; bỏ lỡ chỉ là ngoại lệ chứ không có hệ thống.
- Reading: hiểu gần như mọi văn bản liên quan, gồm lập luận tinh tế, trừu tượng và chi tiết dày; lỗi hiếm và khác biệt ngữ cảnh Academic/GT không gây yếu kém có hệ thống trong cấu trúc tương ứng.
- Writing Task 2 trả lời đề sâu sắc, duy trì lập trường rõ được phát triển hoàn chỉnh và hỗ trợ ý đầy đủ, liên quan. Chất lượng chung mạch lạc tự nhiên, liên kết/chia đoạn tinh tế, vốn từ và cấu trúc hoàn toàn linh hoạt/chính xác, chỉ có lỗi cực hiếm. Academic Task 1 đáp ứng đầy đủ/thích hợp yêu cầu và truyền đạt đặc điểm hình chính yếu chính xác. GT Task 1 hoàn thành đầy đủ mục đích/mọi yêu cầu với văn phong/giọng điệu/bố cục/từ vựng phù hợp quan hệ và chỉ có lỗi ngoại lệ.
- Speaking: nói trôi chảy, mạch lạc, phát triển đầy đủ, từ vựng linh hoạt/chính xác gần như toàn diện, cấu trúc chuẩn xác cao, nối âm linh hoạt, dễ hiểu không gắng sức và chỉ có lỗi mức cao rất hiếm.

## Ràng buộc biến thể và đầu ra

Với tuyên bố Academic Writing Band-N, Task 1 dùng `W-TA-01`, `W-TA-02`, `W-TA-03` cùng Writing leaves chung áp dụng. Với GT, Task 1 dùng `W-GT1-01`, `W-GT1-02`, `W-GT1-03` cùng năng lực Writing chung. Leaf hình `W-TA-*` không bắt buộc ở GT Task 1; leaf thư `W-GT1-*` không bắt buộc ở Academic Task 1. Task 2/chất lượng chung dùng `W-TR-*`, `W-CC-*`, `W-LR-*`, `W-GRA-*` tùy áp dụng.

Ngữ nghĩa đầu ra ở mức kỹ năng:

- Listening Band N: hiệu suất quy đổi thành Band N theo chính sách chấm Listening chính thức áp dụng; Assessment xác định bằng chứng đủ hay chưa.
- Academic Reading Band N: hiệu suất quy đổi theo chính sách Academic Reading và có lấy mẫu đủ cấu trúc đó.
- GT Reading Band N: hiệu suất quy đổi theo chính sách GT và lấy mẫu đủ phân bố ngữ cảnh GT Section 1/2/3; điểm quy đổi/nội dung Academic không thể chứng nhận GT Reading.
- Academic Writing Band N: Academic Task 1 + Task 2 + tiêu chí Writing chung đều đạt ngưỡng Band-N.
- GT Writing Band N: GT Task 1, gồm kiểm soát người nhận/mục đích/điểm yêu cầu/văn phong, + Task 2 + tiêu chí chung đều đạt Band-N. Task 1 và Task 2 không thay thế cho nhau; bằng chứng Task 1 Academic/GT không thể dùng thay chéo.
- Speaking Band N: thể hiện đạt Band-N trên toàn cấu trúc Speaking, chứng minh qua Parts 1, 2, 3; luyện tập một phần không biến tiêu chuẩn Band thành tiêu chuẩn riêng của phần đó.

Đặc tả Band là ngưỡng, không phải danh sách học đầy đủ. Người học Band N có thể còn lỗi được phép ở Band N; đặc điểm của Band cao hơn không được biến thành điều kiện đầu ra bắt buộc của Band N. Điều này giữ ranh giới lỗi còn lại và loại trừ yêu cầu Band cao hơn đã khóa.
