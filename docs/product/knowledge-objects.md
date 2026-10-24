# KNOWLEDGE-OBJECT-ONTOLOGY — Đối tượng kiến thức (Knowledge Objects)

> **THẨM QUYỀN PRODUCT CHÍNH THỨC**
>
> Có thẩm quyền trong domain PRODUCT theo `docs/catalog/project.json`. `CONSTITUTION.md` và `OBJECTIVE.md` giữ thẩm quyền riêng; `contracts/**` có thẩm quyền đối với machine contract chính xác trong phạm vi của nó. Các tham chiếu tới `spec/**` hoặc `design/**` bên dưới, kể cả cách dùng từ “canonical” trong lịch sử, chỉ là nguồn gốc tham khảo, không tạo thẩm quyền ngang hàng.

## Mục đích và quyền sở hữu

Xác định **những điều cần biết** để hỗ trợ các năng lực được giữ trong `docs/product/skill-leaves.md`, có nguồn gốc lịch sử từ `spec/03-SKILLS.md`.

Kiến thức là nội dung ngôn ngữ nền tảng, không phải kỹ năng IELTS thứ năm. Các domain kiến thức chính thức là Grammar (ngữ pháp), Vocabulary (từ vựng) và Phonology (âm vị học).

Shard này lưu giữ kiến thức nền tảng, định danh nguyên tử của Knowledge Object, định nghĩa ngắn gọn, các cạnh tiên quyết từ Knowledge sang Knowledge và ánh xạ từ Skill sang Knowledge.

Danh mục hiện hành có **46 Knowledge Objects nguyên tử với ID ổn định**:

- Grammar: 29;
- Vocabulary: 9;
- Phonology: 8.

## Quy tắc Knowledge Object nguyên tử

Một Knowledge Object thể hiện một khái niệm có thể học độc lập. Tiếp tục tách khi khái niệm vẫn chứa nhiều đơn vị có thể học độc lập. Dừng khi việc tách nhỏ hơn không còn cải thiện học tập, suy luận điều kiện tiên quyết, khắc phục lỗi hoặc tái sử dụng.

## Trường chính thức của Knowledge Object

Mỗi Knowledge Object sở hữu các trường chính thức:

- `id` ổn định;
- tên;
- domain;
- định nghĩa ngắn gọn;
- các cạnh `requires` tường minh trỏ đến Knowledge Objects tiên quyết.

Những mục sau cố ý **không** là trường chính thức của đối tượng, trừ khi bằng chứng trong tương lai nâng cấp chúng:

- ví dụ và bài giải mẫu → nguồn gốc lịch sử `spec/10-CONTENT-MODEL.md` và thẩm quyền DATA/nội dung hiện hành;
- danh mục nhận thức sai/cách khắc phục → nguồn gốc lịch sử `spec/10-CONTENT-MODEL.md` và thẩm quyền DATA/nội dung hiện hành;
- liên hệ mềm `related_to` không ảnh hưởng cổng điều kiện/lựa chọn → chỉ thuộc đồ thị lịch sử/hỗ trợ;
- mức Band liên quan → các shard PRODUCT về IELTS/Bands và Curriculum;
- trạng thái tiếp thu của người học → thẩm quyền DATA/BEHAVIOR hiện hành.

## Đồ thị ngữ pháp (Grammar)

| ID | Đối tượng | Định nghĩa chính thức | Yêu cầu tiên quyết |
|---|---|---|---|
| `K-GRA-010` | Loại từ / từ loại | Danh từ, động từ, tính từ, trạng từ, giới từ, liên từ, từ hạn định, đại từ và các nhóm chức năng liên quan. | — |
| `K-GRA-061` | Danh từ đếm được và số nhiều | Cách dùng danh từ đếm được/không đếm được, số nhiều có/không theo quy tắc và đánh dấu số lượng. | `K-GRA-010` |
| `K-GRA-001` | Cấu trúc mệnh đề | Thành phần cốt lõi như chủ ngữ, động từ, tân ngữ, bổ ngữ, trạng ngữ và trật tự cơ bản. | `K-GRA-010` |
| `K-GRA-002` | Câu đơn | Câu hoàn chỉnh gồm một mệnh đề độc lập. | `K-GRA-001` |
| `K-GRA-003` | Câu ghép / phối hợp | Từ hai mệnh đề độc lập trở lên được nối bằng phép phối hợp. | `K-GRA-002`, `K-GRA-020` |
| `K-GRA-004` | Câu phức / phụ thuộc | Mệnh đề độc lập kết hợp với một hoặc nhiều mệnh đề phụ. | `K-GRA-002`, `K-GRA-021` |
| `K-GRA-005` | Mệnh đề quan hệ | Mệnh đề bổ nghĩa danh từ bằng dạng quan hệ, gồm dạng xác định và không xác định. | `K-GRA-004`, `K-GRA-031` |
| `K-GRA-006` | Mệnh đề điều kiện | Quan hệ điều kiện biểu đạt qua mẫu câu if/điều kiện, gồm cả các dạng hỗn hợp thông dụng. | `K-GRA-004` |
| `K-GRA-007` | Mệnh đề danh từ | Mệnh đề phụ hoạt động như thành phần danh từ. | `K-GRA-004` |
| `K-GRA-020` | Liên từ đẳng lập | Các từ nối như and/but/or/so liên kết những đơn vị ngữ pháp ngang hàng. | `K-GRA-010` |
| `K-GRA-021` | Liên từ phụ thuộc | Các từ nối như although/because/if/while mở đầu mệnh đề phụ thuộc. | `K-GRA-010` |
| `K-GRA-022` | Trạng ngữ liên kết / từ nối | Chuyển ý trong câu/diễn ngôn như however, therefore, moreover và consequently. | `K-GRA-010` |
| `K-GRA-030` | Đại từ | Hệ thống đại từ nhân xưng, sở hữu, chỉ định, phản thân và các loại liên quan. | `K-GRA-010` |
| `K-GRA-031` | Đại từ quan hệ | Các dạng quan hệ như who, whom, which, that và whose. | `K-GRA-030` |
| `K-GRA-032` | Từ hạn định | Mạo từ, từ chỉ định, lượng từ, từ sở hữu và thành phần giới hạn cụm danh từ khác. | `K-GRA-010` |
| `K-GRA-033` | Quy chiếu và thay thế tạo liên kết | Phương tiện quy chiếu/thay thế liên kết diễn ngôn và tránh lặp lại không cần thiết. | `K-GRA-030`, `K-GRA-032` |
| `K-GRA-040` | Mạo từ: a / an / the | Các dạng mạo từ xác định, không xác định và cách dùng cơ bản. | `K-GRA-032` |
| `K-GRA-041` | Quy tắc mạo từ xác định / không xác định / không mạo từ | Điều kiện lựa chọn mạo từ, gồm không dùng mạo từ và quy chiếu chung/riêng. | `K-GRA-040` |
| `K-GRA-050` | Thì hiện tại đơn | Dạng hiện tại đơn và cách dùng chính cho sự thật, thói quen, trạng thái và khái quát. | `K-GRA-002` |
| `K-GRA-051` | Thì quá khứ đơn | Dạng quá khứ đơn và cách dùng chính cho sự kiện/trạng thái đã hoàn tất trong quá khứ. | `K-GRA-002` |
| `K-GRA-052` | Thì hiện tại hoàn thành | Dạng hiện tại hoàn thành nối sự kiện/trạng thái trước đó với ý nghĩa hiện tại. | `K-GRA-051` |
| `K-GRA-053` | Thì quá khứ hoàn thành | Dạng quá khứ hoàn thành cho sự kiện/trạng thái trước một mốc quá khứ khác. | `K-GRA-052` |
| `K-GRA-054` | Thể tiếp diễn | Thể `be + -ing` cho tình huống đang diễn ra hoặc bị giới hạn thời gian tại những mốc khác nhau. | `K-GRA-050`, `K-GRA-051` |
| `K-GRA-055` | Dạng chỉ tương lai | Các dạng chỉ tương lai thông dụng như will, going to và mẫu hiện tại tiếp diễn. | `K-GRA-002` |
| `K-GRA-060` | Hòa hợp chủ ngữ–động từ | Dạng động từ hữu hạn phù hợp với ngôi/số của chủ ngữ, gồm tình huống phức tạp thường gặp. | `K-GRA-002`, `K-GRA-061` |
| `K-GRA-062` | Động từ khiếm khuyết | Trợ động từ như can, could, must, should, may, might và would cùng nghĩa cốt lõi. | `K-GRA-002` |
| `K-GRA-063` | Thể bị động | Cấu trúc bị động với dạng `be` phù hợp và quá khứ phân từ để giảm nhấn tác nhân/tập trung vào đối tượng chịu tác động. | `K-GRA-002`, `K-GRA-061` |
| `K-GRA-064` | So sánh hơn và so sánh nhất | Hình thái/cú pháp so sánh hơn và nhất của tính từ, trạng từ. | `K-GRA-010` |
| `K-GRA-065` | Phủ định và dạng câu hỏi | Phủ định cơ bản, câu hỏi yes/no và wh-, gồm cách dùng trợ động từ và trật tự từ. | `K-GRA-002` |

## Đồ thị từ vựng (Vocabulary)

| ID | Đối tượng | Định nghĩa chính thức | Yêu cầu tiên quyết |
|---|---|---|---|
| `K-VOC-010` | Từ vựng cốt lõi / tần suất cao | Các họ từ tiếng Anh thông dụng cần cho hiểu và diễn đạt cơ bản. | — |
| `K-VOC-011` | Từ vựng học thuật | Vốn từ trang trọng/học thuật thường gặp trong đọc và viết học thuật. | `K-VOC-010` |
| `K-VOC-012` | Bộ từ theo chủ đề | Các bộ từ có thể tái sử dụng, tổ chức theo chủ đề thường liên quan IELTS. | `K-VOC-010` |
| `K-VOC-020` | Collocations (kết hợp từ) | Các tổ hợp từ thường đi với nhau theo quy ước và giới hạn cách dùng. | `K-VOC-010` |
| `K-VOC-021` | Thành ngữ và cụm cố định | Cụm nhiều từ cố định/bán cố định với nghĩa hoặc cách dùng theo quy ước. | `K-VOC-010` |
| `K-VOC-030` | Cấu tạo từ / phụ tố | Quan hệ tạo dạng giữa các từ cùng họ qua tiền tố, hậu tố và phái sinh. | `K-VOC-010` |
| `K-VOC-031` | Quy tắc và mẫu chính tả | Quy ước chính tả tiếng Anh và mẫu lặp lại cần cho viết chính xác. | `K-VOC-010` |
| `K-VOC-040` | Sắc thái và mức độ trang trọng | Lựa chọn từ theo văn phong trang trọng, trung tính, thân mật và ngữ cảnh. | `K-VOC-010` |
| `K-VOC-041` | Từ đồng nghĩa và diễn đạt lại | Các cách dùng từ thay thế, sắc thái nghĩa và giới hạn sử dụng để diễn đạt chính xác cùng ý. | `K-VOC-010` |

`K-VOC-012` đại diện cho hệ thống bộ từ theo chủ đề. Danh sách từ của từng chủ đề là nội dung cụ thể theo thẩm quyền DATA/nội dung hiện hành, không phải domain chính thức mới.

Từ, cụm từ, collocation, ví dụ và thẻ ôn tập cá nhân người học lưu cũng là các thể hiện nội dung/ôn tập cụ thể. Chúng có thể tham chiếu một hoặc nhiều Knowledge Objects chính thức, nhưng hành động lưu không tạo ID `K-*` mới.

## Đồ thị âm vị học (Phonology)

| ID | Đối tượng | Định nghĩa chính thức | Yêu cầu tiên quyết |
|---|---|---|---|
| `K-PHON-010` | Âm vị phụ âm | Tập âm phụ âm tiếng Anh và phân biệt cấu âm cần cho nói/nghe dễ hiểu. | — |
| `K-PHON-011` | Âm vị nguyên âm | Hệ thống nguyên âm đơn/đôi tiếng Anh và các đối lập nguyên âm có ý nghĩa. | — |
| `K-PHON-012` | Đối lập âm vị và cặp tối thiểu | Các cặp âm đối lập giúp nhận biết và tạo âm vị phân biệt nghĩa. | `K-PHON-010`, `K-PHON-011` |
| `K-PHON-020` | Trọng âm từ | Vị trí trọng âm chính và biến đổi trọng âm phổ biến giữa các từ cùng họ. | `K-PHON-011` |
| `K-PHON-021` | Trọng âm câu | Mẫu nhấn nổi bật trong phát ngôn, gồm nhấn từ nội dung/từ chức năng. | `K-PHON-020` |
| `K-PHON-030` | Ngữ điệu | Đường cao độ và chức năng diễn ngôn, ngữ dụng hoặc thái độ. | `K-PHON-021` |
| `K-PHON-040` | Nối âm trong lời nói | Nối âm, giảm âm, đồng hóa, lược âm và dạng yếu qua ranh giới từ. | `K-PHON-010`, `K-PHON-011` |
| `K-PHON-041` | Nhịp điệu và chia cụm | Nhóm lời nói thành đơn vị nghĩa và giữ tổ chức nhịp điệu dễ hiểu. | `K-PHON-021`, `K-PHON-040` |

Đồ thị âm vị học độc lập với L1 (ngôn ngữ thứ nhất). Ưu tiên đối lập âm theo L1 thuộc nội dung bản địa hóa/khắc phục, không thuộc đồ thị này.

## Phân giải Skill→Knowledge

Đây là thẩm quyền PRODUCT duy nhất cho ánh xạ nhu cầu kiến thức từ Skill Leaf sang Knowledge Objects. `spec/04-KNOWLEDGE.md` chỉ là nguồn gốc lịch sử.

Một cạnh trong ánh xạ nghĩa là Knowledge Object là **điều kiện tiên quyết nội tại, phổ quát của Skill Leaf trong toàn bộ phạm vi chính thức áp dụng cho leaf đó**. Vì vậy cạnh có thể trở thành điều kiện Required ở bước sau. Kiến thức chỉ hữu ích với một biến thể, chủ đề, kiểu trình bày, prompt, stimulus hay nội dung cụ thể **không** thuộc ánh xạ phổ quát này; Curriculum/Content có thể chọn tường minh khi cần.

### Writing

| Skill Leaf | Knowledge Objects được phân giải |
|---|---|
| `W-CC-03` | `K-GRA-020`, `K-GRA-021`, `K-GRA-022` |
| `W-CC-04` | `K-GRA-030`, `K-GRA-032`, `K-GRA-033` |
| `W-LR-01` | `K-VOC-012` |
| `W-LR-02` | `K-VOC-020` |
| `W-LR-03` | `K-VOC-041`, `K-VOC-020` |
| `W-LR-05` | `K-VOC-030`, `K-VOC-031` |
| `W-GRA-01` | `K-GRA-001`, `K-GRA-002` |
| `W-GRA-02` | `K-GRA-003`, `K-GRA-020` |
| `W-GRA-03` | `K-GRA-004`, `K-GRA-021`, `K-GRA-005`, `K-GRA-006`, `K-GRA-007`, `K-GRA-031` |
| `W-GRA-04` | `K-GRA-050`, `K-GRA-051`, `K-GRA-052`, `K-GRA-053`, `K-GRA-054`, `K-GRA-055` |
| `W-GRA-05` | `K-GRA-032`, `K-GRA-040`, `K-GRA-041` |
| `W-GRA-07` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007`, `K-GRA-062`, `K-GRA-063` |
| `W-GT1-02` | `K-VOC-040` |

`W-GT1-01` và `W-GT1-03` là leaf năng lực thực hiện nhiệm vụ; nhu cầu kiến thức trực tiếp đã được bao phủ bởi các mục tiêu từ vựng/ngữ pháp/liên kết chung được chọn cho lá thư cụ thể. Không cần Knowledge Object riêng cho thư.

`W-TA-03` không có một điều kiện tiên quyết Knowledge đơn lẻ phổ quát ngoài kiến thức Writing chung. So sánh (`K-GRA-064`), thì/thể, ngôn ngữ số liệu, quy trình và không gian được chọn khi kiểu trình bày/stimulus cụ thể của Academic Task 1 khiến chúng cần thiết; không được chặn cứng bài process/map chỉ vì dạng Task 1 khác sử dụng nhiều so sánh.

### Speaking

| Skill Leaf | Knowledge Objects được phân giải |
|---|---|
| `S-FC-03` | `K-GRA-020`, `K-GRA-021`, `K-GRA-022` |
| `S-LR-01` | `K-VOC-012` |
| `S-LR-04` | `K-VOC-040`, `K-VOC-020` |
| `S-GRA-01` | `K-GRA-001`, `K-GRA-002` |
| `S-GRA-02` | `K-GRA-004`, `K-GRA-021`, `K-GRA-005` |
| `S-GRA-04` | `K-GRA-001`, `K-GRA-002`, `K-GRA-060`, `K-GRA-065` |
| `S-P-01` | `K-PHON-010`, `K-PHON-011`, `K-PHON-012` |
| `S-P-02` | `K-PHON-020` |
| `S-P-03` | `K-PHON-021`, `K-PHON-030` |

### Reading

Chưa có Reading Skill Leaf nào cần thêm điều kiện tiên quyết Knowledge **phổ quát, không phụ thuộc biến thể** trong ánh xạ hiện tại.

Đặc biệt, `R-COMP-06` dùng chung không bắt buộc phổ quát `K-VOC-011` Academic vocabulary: Academic và General Training dùng chung leaf này nhưng vốn từ corpus/ngữ cảnh khác nhau. Lộ trình Academic có thể xếp lịch `K-VOC-011` tường minh; GT dùng nội dung từ vựng chung/theo chủ đề/nơi làm việc và phân bố ngữ cảnh phù hợp, không kế thừa cổng cứng từ vựng Academic ẩn.

Độ đa dạng ngữ cảnh Reading của General Training là điều kiện curriculum/content/evidence, không phải Knowledge Object riêng.

Listening hiện không có cạnh điều kiện Knowledge tường minh trong đồ thị đã đóng băng; các điều kiện trực tiếp là cạnh năng lực trong cùng kỹ năng ở `docs/product/skill-leaves.md`.

## Ngữ nghĩa phụ thuộc

Shard này giữ các cạnh điều kiện tiên quyết Knowledge phổ quát. Thẩm quyền Curriculum PRODUCT có thể thêm mục tiêu học theo ngữ cảnh/biến thể và thứ tự node Recommended tường minh, nhưng không được diễn giải Knowledge Object chỉ hữu ích trong một ngữ cảnh thành cạnh Required nội tại cho mọi trường hợp của Skill Leaf dùng chung.

Ngữ nghĩa cổng Required / Recommended / Independent được sử dụng bởi thẩm quyền Curriculum PRODUCT và được thực thi theo ngữ nghĩa progression của BEHAVIOR/DATA.

Định danh đối tượng và ngữ nghĩa điều kiện tiên quyết là chính thức. Độ khó, thời gian học, tần suất, ví dụ, hiểu lầm phổ biến, ưu tiên bản địa hóa, bộ từ theo ngữ cảnh và danh mục khắc phục thuộc vấn đề thực nghiệm/nội dung, trừ khi bằng chứng tương lai nâng cấp tường minh chúng.
