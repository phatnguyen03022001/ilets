# CURRICULUM-COVERAGE Ngữ nghĩa chương trình học và hỗ trợ sản phẩm

> **THẨM QUYỀN PRODUCT CHÍNH THỨC**
>
> Có thẩm quyền trong domain PRODUCT theo `docs/catalog/project.json`. `CONSTITUTION.md` và `OBJECTIVE.md` giữ thẩm quyền riêng, còn `contracts/**` giữ thẩm quyền machine contract chính xác trong phạm vi của nó. Tham chiếu tới `spec/**` hoặc `design/**` bên dưới, kể cả cách gọi “canonical” trong lịch sử, chỉ để truy nguyên nguồn gốc, không tạo thẩm quyền ngang hàng.

## Ngữ nghĩa chương trình học

Chương trình học hiện hành có **44 Curriculum Nodes `C-*` với định danh ổn định** xuyên suốt Band 3–9. Mỗi node sở hữu ID/giai đoạn Band ổn định, vị trí trình tự khuyến nghị, các mục tiêu Skill/Knowledge, trọng tâm, kết quả mong đợi, các phụ thuộc node được khuyến nghị tường minh và mục đích đầu ra node. Ước lượng giờ/tải học tập cũ không phải dữ liệu chính thức. Hoàn thành node không đồng nghĩa chứng nhận Band.

Các lớp phụ thuộc có ý nghĩa chính xác:

- **Required** — điều kiện tiên quyết nội tại Skill→Skill, Knowledge→Knowledge hoặc Skill→Knowledge. Cạnh Required chưa thỏa sẽ chặn cứng việc học phụ thuộc. Các cạnh đối tượng Required kế thừa từ owner Skill/Knowledge.
- **Recommended** — thứ tự Curriculum Node hữu ích, mã hóa bằng `Depends`; không phải cổng chặn cứng khi bằng chứng cho thấy lộ trình khác phù hợp.
- **Independent** — không có quan hệ tiên quyết.

`Depends` chỉ chứa ID `C-*` ổn định và biểu thị trình tự Recommended. Không được chứa ID Skill/Knowledge hoặc tên mô tả. Văn bản trọng tâm/kết quả dễ đọc không tạo phụ thuộc ngầm.

### Hợp đồng selector xác định

Chỉ các token selector lúc biên dịch sau được phép xuất hiện trong cột mục tiêu node:

```text
@task1-analysis
@task1-fulfilment
@task1-global-control
@variant-reading-advanced-knowledge
@band-phase-skill-set
@band-phase-knowledge-set
```

`@task1-*` và `@variant-reading-advanced-knowledge` được phân giải theo biến thể Academic/GT đã chọn. `@band-phase-skill-set` là hợp của mọi mục tiêu Skill đã phân giải đầy đủ trong các node trước đó cùng giai đoạn Band; `@band-phase-knowledge-set` là hợp Knowledge tường minh tương ứng. Mở rộng đệ quy đến khi chỉ còn ID chính thức. Selector chưa phân giải tại ranh giới materialization là không hợp lệ; biến thể thực sự chưa xác định chỉ cho phép node dùng chung không cần biến thể còn thiếu. Không được tự chế selector.

Phân giải selector theo biến thể:

| Selector | Academic | General Training |
|---|---|---|
| `@task1-analysis` | `W-TA-01` | `W-GT1-01` |
| `@task1-fulfilment` | `W-TA-03` | `W-GT1-03` |
| `@task1-global-control` | `W-TA-02` | `W-GT1-02` |
| `@variant-reading-advanced-knowledge` | `K-VOC-011` | tập rỗng |

Phân giải thành tập rỗng với GT nghĩa là không có một Knowledge Object bổ sung nào bắt buộc phổ quát chỉ vì bài Reading GT dày thông tin. Khác biệt Reading theo biến thể vẫn là điều kiện ngữ cảnh/nội dung/readiness. Academic sử dụng bài đọc/corpus Academic; GT giữ phân bố Section 1 đời sống thường ngày, Section 2 nơi làm việc và Section 3 chủ đề phổ thông dài hơn. Người học GT không cần leaf Academic `W-TA-*` cho GT Task 1, và người học Academic không cần `W-GT1-*`; năng lực dùng chung vẫn tái sử dụng được.

### Danh mục Curriculum Node ổn định

| Node | Trọng tâm | Mục tiêu Skill | Mục tiêu Knowledge | Depends được khuyến nghị | Mục đích đầu ra |
|---|---|---|---|---|---|
| `C-B3-01` | Từ loại, cấu trúc mệnh đề, câu đơn | `W-GRA-01`, `S-GRA-01` | `K-GRA-010`, `K-GRA-001`, `K-GRA-002` | — | lõi câu viết/nói cơ bản chính xác; tiếp thu ngữ pháp nền |
| `C-B3-02` | Danh từ, đếm được, từ hạn định/mạo từ | `W-GRA-05` | `K-GRA-061`, `K-GRA-032`, `K-GRA-040` | `C-B3-01` | kiểm soát cơ bản danh từ đếm được/từ hạn định/mạo từ |
| `C-B3-03` | Thì cơ bản và hòa hợp | `W-GRA-04` | `K-GRA-050`, `K-GRA-051`, `K-GRA-060` | `C-B3-01` | hiện tại/quá khứ cơ bản và hòa hợp trong diễn đạt đơn giản |
| `C-B3-04` | Từ vựng cốt lõi và chính tả | `W-LR-01`, `W-LR-05`, `S-LR-01` | `K-VOC-010`, `K-VOC-031` | — | từ vựng tần suất cao, kiểm soát chính tả cơ bản |
| `C-B3-05` | Nền âm vị và độ dễ hiểu | `S-P-01`, `S-P-05` | `K-PHON-010`, `K-PHON-011`, `K-PHON-012` | — | phân biệt âm cốt lõi với độ dễ hiểu rộng |
| `C-B3-06` | Nắm ý chính và chi tiết khi tiếp nhận | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-010` | `C-B3-04` | nắm ý chính/chi tiết quan trọng từ đầu vào nói/viết đơn giản |
| `C-B3-07` | Dạng câu hỏi tiếp nhận cơ bản | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B3-06` | làm câu tiếp nhận cơ bản trong giới hạn định dạng/số từ |
| `C-B3-08` | Tích hợp nền tảng | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B3-01`, `C-B3-02`, `C-B3-03`, `C-B3-04`, `C-B3-05`, `C-B3-06`, `C-B3-07` | tích hợp Band 3; không phải cổng chứng nhận bốn kỹ năng |
| `C-B4-01` | Chính xác Writing, dấu câu, chia đoạn | `W-GRA-01`, `W-GRA-06`, `W-CC-01` | — | `C-B3-01`, `C-B3-03` | viết đoạn đơn giản với dấu câu |
| `C-B4-02` | Phân tích Task 1 và đề Task 2 | `@task1-analysis`, `W-TR-01` | — | `C-B3-08` | yêu cầu Task 1 theo biến thể đã chọn và phân tích đề Task 2 |
| `C-B4-03` | Từ vựng chủ đề và chọn từ | `W-LR-01`, `S-LR-01`, `S-LR-04` | `K-VOC-012`, `K-VOC-040` | `C-B3-04` | đủ từ theo chủ đề quen thuộc, đúng ngữ cảnh |
| `C-B4-04` | Duy trì nói và độ chính xác cơ bản | `S-FC-01`, `S-GRA-04`, `S-P-05` | `K-GRA-010` | `C-B3-01`, `C-B3-05` | nói liên tục cơ bản về chủ đề quen thuộc |
| `C-B4-05` | Ý chính và chi tiết khi tiếp nhận | `L-COMP-01`, `L-COMP-02`, `R-COMP-01`, `R-COMP-02` | `K-VOC-012` | `C-B3-06` | nắm ý chính/chi tiết từ đầu vào phức tạp hơn vừa phải |
| `C-B4-06` | Dạng tiếp nhận cơ bản ở độ khó Band 4 | `L-QT-01`, `L-QT-05`, `R-QT-05` | — | `C-B4-05` | chiến lược tiếp nhận cơ bản phù hợp độ khó hiện tại |
| `C-B4-07` | Tích hợp Band 4 | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B4-01`, `C-B4-02`, `C-B4-03`, `C-B4-04`, `C-B4-05`, `C-B4-06` | mốc điều phối Band 4, không chứng nhận |
| `C-B5-01` | Nền câu ghép/câu phức | `W-GRA-02`, `W-GRA-03` | `K-GRA-003`, `K-GRA-020`, `K-GRA-004`, `K-GRA-021` | `C-B4-01` | sử dụng được câu ghép/câu phức |
| `C-B5-02` | Liên kết và tiến triển logic | `W-CC-02`, `W-CC-03`, `W-CC-04` | `K-GRA-022`, `K-GRA-033` | `C-B4-01`, `C-B5-01` | tiến triển, nối ý, quy chiếu/thay thế |
| `C-B5-03` | Lập trường, phát triển ý, bám đề Task 2 | `W-TR-02`, `W-TR-03`, `W-TR-04` | `K-VOC-012` | `C-B4-02`, `C-B5-02` | trả lời Task 2 đúng trọng tâm với lập trường có hỗ trợ rõ |
| `C-B5-04` | Đáp ứng nội dung Task 1 | `@task1-fulfilment` | — | `C-B4-02` | đáp ứng nội dung Task 1 theo biến thể đã chọn |
| `C-B5-05` | Diễn đạt lại, cấu tạo từ, chính tả | `W-LR-03`, `W-LR-05` | `K-VOC-030`, `K-VOC-031` | `C-B4-03` | diễn đạt lại giữ nghĩa, dùng đúng dạng từ/chính tả |
| `C-B5-06` | Trôi chảy, độ phức tạp và phát âm Speaking | `S-FC-02`, `S-FC-03`, `S-FC-04`, `S-FC-05`, `S-GRA-02`, `S-LR-02`, `S-P-02`, `S-P-03` | `K-GRA-021`, `K-VOC-012` | `C-B4-04`, `C-B5-01` | lượt nói dài mạch lạc, ngôn ngữ/trọng âm/ngữ điệu phức tạp hơn |
| `C-B5-07` | Suy luận, diễn đạt lại, gây nhiễu khi tiếp nhận | `L-COMP-03`, `L-COMP-04`, `L-COMP-05`, `R-COMP-03`, `R-COMP-04`, `R-COMP-05` | — | `C-B4-05` | suy luận/diễn đạt lại/cấu trúc/gây nhiễu đáng tin cậy hơn |
| `C-B5-08` | Dạng câu hỏi tiếp nhận bậc cao | `L-QT-02`, `L-QT-03`, `L-QT-04`, `R-QT-01`, `R-QT-02`, `R-QT-04`, `R-QT-06` | — | `C-B5-07` | chiến lược dạng câu hỏi bậc cao dưới giới hạn thực tế |
| `C-B5-09` | Tích hợp Band 5 | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B5-01`, `C-B5-02`, `C-B5-03`, `C-B5-04`, `C-B5-05`, `C-B5-06`, `C-B5-07`, `C-B5-08` | tích hợp Band 5, không chứng nhận |
| `C-B6-01` | Chính xác và đa dạng ngữ pháp | `W-GRA-03`, `W-GRA-02`, `S-GRA-03` | `K-GRA-004`, `K-GRA-003` | `C-B5-01` | tăng chính xác/đa dạng khi diễn đạt |
| `C-B6-02` | Kiểm soát tổng thể Task 1 và quy chiếu linh hoạt | `@task1-global-control`, `W-CC-04` | — | `C-B5-02`, `C-B5-04` | kiểm soát Task 1 toàn cục theo biến thể, giảm lặp |
| `C-B6-03` | Collocation và thành ngữ | `W-LR-02`, `S-LR-03` | `K-VOC-020`, `K-VOC-021` | `C-B4-03` | collocation chính xác và bước đầu dùng từ hiếm/thành ngữ phù hợp |
| `C-B6-04` | Tiếp nhận nội dung dài/dày và chuyển giao biến thể | `L-COMP-06`, `R-COMP-06` | `@variant-reading-advanced-knowledge` | `C-B5-07` | xử lý đầu vào dài/dày và chuyển giao ngữ cảnh biến thể |
| `C-B6-05` | Phân loại quan điểm tác giả | `R-QT-03` | — | `C-B5-08` | phân biệt tác giả đồng tình/không đồng tình/không đề cập |
| `C-B6-06` | Tích hợp Band 6 | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B6-01`, `C-B6-02`, `C-B6-03`, `C-B6-04`, `C-B6-05` | tích hợp Band 6, không chứng nhận |
| `C-B7-01` | Linh hoạt cấu trúc / tần suất đúng | `W-GRA-07`, `W-GRA-03` | `K-GRA-005`, `K-GRA-006`, `K-GRA-007` | `C-B6-01` | cấu trúc phức linh hoạt/đa dạng với độ chính xác thường xuyên |
| `C-B7-02` | Từ vựng nâng cao | `W-LR-04`, `S-LR-03` | `K-VOC-021` | `C-B6-03` | vốn từ ít thông dụng/thành ngữ ngày càng chính xác |
| `C-B7-03` | Liên kết, lập trường, phát triển ý | `W-CC-03`, `W-CC-04`, `W-TR-02`, `W-TR-03` | `K-GRA-022`, `K-GRA-033` | `C-B5-03`, `C-B6-02` | duy trì lập trường và liên kết rõ, phát triển đầy đủ |
| `C-B7-04` | Suy luận và cấu trúc khi tiếp nhận | `R-COMP-04`, `R-COMP-05`, `L-COMP-03`, `L-COMP-04` | — | `C-B6-04` | suy luận/lập trường/diễn đạt lại/cấu trúc đáng tin cậy |
| `C-B7-05` | Phát âm đa dạng và linh hoạt | `S-P-03`, `S-P-04` | `K-PHON-030`, `K-PHON-040` | `C-B5-06` | ngữ điệu/nối âm/chia cụm hữu ích trong lượt nói dài |
| `C-B7-06` | Tích hợp Band 7 | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B7-01`, `C-B7-02`, `C-B7-03`, `C-B7-04`, `C-B7-05` | tích hợp Band 7, không chứng nhận |
| `C-B8-01` | Ngôn ngữ rộng/linh hoạt, gần như không lỗi | `W-GRA-07`, `W-LR-04`, `W-GRA-01`, `W-GRA-03`, `S-GRA-03` | `K-GRA-005`, `K-GRA-007` | `C-B7-01`, `C-B7-02` | ngôn ngữ đa dạng, linh hoạt, chính xác cao |
| `C-B8-02` | Liên kết khéo léo/phản hồi phát triển đủ | `W-CC-02`, `W-CC-03`, `W-CC-04`, `W-TR-03`, `@task1-global-control` | `K-GRA-022`, `K-GRA-033` | `C-B7-03` | bài trả lời phát triển tốt, dễ theo dõi, gồm kiểm soát Task 1 theo biến thể |
| `C-B8-03` | Speaking trôi chảy tự nhiên/vốn từ rộng | `S-FC-02`, `S-FC-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-PHON-040`, `K-PHON-041` | `C-B7-05` | lời nói trôi chảy, giàu ngôn ngữ, dễ hiểu |
| `C-B8-04` | Hiểu nội dung tiếp nhận phức tạp | `L-COMP-06`, `R-COMP-06`, `R-COMP-04` | — | `C-B7-04` | ít khó khăn với nội dung phức tạp/trừu tượng ở nhiều ngữ cảnh biến thể |
| `C-B8-05` | Tích hợp Band 8 | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B8-01`, `C-B8-02`, `C-B8-03`, `C-B8-04` | tích hợp Band 8, không chứng nhận |
| `C-B9-01` | Linh hoạt và chính xác toàn diện | `W-GRA-07`, `W-LR-04`, `S-GRA-03`, `S-LR-03`, `S-P-05` | `K-GRA-005`, `K-GRA-007`, `K-VOC-021` | `C-B8-01`, `C-B8-02`, `C-B8-03` | gần như toàn diện về linh hoạt/chính xác/độ dễ hiểu bền vững |
| `C-B9-02` | Thành thạo tổng hợp ở mức bài thi | `@task1-global-control`, `W-TR-03`, `W-CC-03`, `R-COMP-04`, `L-COMP-06` | — | `C-B8-05` | đáp ứng yêu cầu mô phỏng bài thi độc lập mức trần theo biến thể |
| `C-B9-03` | Tích hợp mức trần | `@band-phase-skill-set` | `@band-phase-knowledge-set` | `C-B9-01`, `C-B9-02` | tích hợp cuối; chứng nhận vẫn thuộc Assessment/Progression |

Trình tự thích ứng chỉ được đổi thứ tự/xen kẽ node khi điều kiện Required, đầy đủ kết quả mục tiêu và yêu cầu biến thể vẫn được đáp ứng. Người học có thể ở giai đoạn Band khác nhau theo từng kỹ năng. Node tích hợp không đồng bộ chứng nhận bốn kỹ năng. Materialization tĩnh phải thất bại nếu có ID/selector lạ, selector chưa phân giải khi cần, văn bản mô tả đặt ở vị trí ID, node `Depends` thiếu, chu trình Recommended hoặc mở rộng Task 1 sai biến thể.

## Ngữ nghĩa hỗ trợ sản phẩm

Mục tiêu sản phẩm trong phạm vi xác định đi theo đúng chuỗi trạng thái:

```text
MODELLED → COVERED → SUPPORTED_FOR_PRODUCT → VALIDATED
```

- `MODELLED`: cấu trúc/ngữ nghĩa sản phẩm được biểu diễn đủ để mở rộng điều kiện áp dụng và phơi bày khoảng trống; chưa phải hỗ trợ thực thi.
- `COVERED`: lộ trình chạy được hoàn chỉnh, không có CoverageGap chặn, và mọi điều kiện TargetCoverageSpecification áp dụng ngoài `validation` đều `SATISFIED`. Bất kỳ điều kiện không phải validation có trạng thái `UNKNOWN`, `DEFINED`, `PARTIAL`, `BLOCKED` hoặc `CALIBRATION_REQUIRED` sẽ chặn COVERED.
- `SUPPORTED_FOR_PRODUCT`: TargetSupportDeclaration có phiên bản kích hoạt mục tiêu đã COVERED cho ranh giới sản phẩm/release cụ thể, xác nhận phạm vi/phiên bản chính xác và các cổng cần cho release. Văn bản tuyên bố không được vượt qua COVERED.
- `VALIDATED`: trạng thái SUPPORTED_FOR_PRODUCT hiện hành cộng `validation = SATISFIED` thực nghiệm trong phạm vi tên phiên bản/điều kiện người học, sản phẩm, nội dung, evaluator và can thiệp. Kiến trúc nhất quán không phải bằng chứng validation.

Từ vựng trạng thái điều kiện bao phủ là hệ riêng:

```text
UNKNOWN
DEFINED
PARTIAL
SATISFIED
BLOCKED
NOT_APPLICABLE
CALIBRATION_REQUIRED
```

`UNKNOWN` không bao giờ đạt; `DEFINED` nghĩa là đã biết tiêu chí nhưng chưa hiện thực/kiểm chứng thực thi đủ; `PARTIAL` nghĩa là vài điều kiện con đạt nhưng phần còn lại thiếu/chưa kiểm; `SATISFIED` nghĩa là mọi điều kiện con áp dụng đều đạt; `BLOCKED` là lỗi cứng/thiếu tiên quyết đã biết; `NOT_APPLICABLE` phải do owner suy ra không áp dụng; `CALIBRATION_REQUIRED` nghĩa là đã có cơ chế/lộ trình nhưng chưa hiệu chuẩn đủ cho hậu quả. Thứ tự tổng hợp phải giữ điều kiện áp dụng chưa rõ, trở ngại cứng và hiệu chuẩn thay vì lấy trung bình trạng thái. `OUT_OF_SCOPE` là quyết định phạm vi, không phải trạng thái điều kiện/hỗ trợ sản phẩm.

### Phân loại CoverageGap

CoverageGap là năng lực sản phẩm thiếu hoặc điều kiện sản phẩm chưa giải quyết, **không phải điểm yếu người học**. Các lớp chính thức:

`MODEL_OR_SPEC`, `INTERVENTION_OR_ACTIVITY`, `CONTENT_OR_ASSET`, `EVIDENCE_OR_EVALUATOR`, `EXPERIENCE`, `TRANSITION`, `CONTRACT_OR_INTEGRATION`, `COST_OR_OPERATIONS`, `RIGHTS_PRIVACY_RELIABILITY`, `CALIBRATION_OR_VALIDATION`.

Mỗi gap giữ mục tiêu/điều kiện trong phạm vi, điều kiện thiếu/thất bại, hậu quả chặn, phụ thuộc, lớp nhu cầu và nguồn gốc/phiên bản.

### Các điều kiện TargetCoverageSpecification bắt buộc

Mọi mục tiêu trong phạm vi phải đánh giá độc lập từng điều kiện áp dụng:

| Điều kiện | Ý nghĩa bắt buộc |
|---|---|
| `construct_model` | biểu diễn chính xác cấu trúc IELTS bên ngoài hiện hành |
| `official_family_coverage` | mọi ID dạng thi chính thức áp dụng đều có đường thực thi sản phẩm/nội dung/bằng chứng |
| `material_subformat_coverage` | các Presentation Classes bắt buộc được biểu diễn khi một nhãn dạng thi che giấu khoảng trống đáng kể |
| `skill_knowledge_model` | năng lực/kiến thức chính thức cần thiết tồn tại |
| `band_threshold` | xác định ngưỡng Band/nhiệm vụ/biến thể mục tiêu |
| `curriculum_route` | tồn tại lộ trình theo thứ tự/biến thể hợp lệ |
| `practice_intervention` | learning mechanism/mode bắt buộc chạy được |
| `feature_experience` | người học thực hiện tương tác, hiểu trạng thái/kết quả; hành vi runtime/browser bắt buộc sử dụng được |
| `content_assets` | đủ nội dung chạy được hợp lệ cho dạng/ngữ cảnh/đa dạng; chỉ tính generator khi release phụ thuộc nó |
| `assessment_policy` | có Assessment Type phù hợp và EvidenceRequirement thực thi được, có phiên bản |
| `evaluator_scoring` | có chấm điểm xác định hoặc đánh giá productive đã hiệu chuẩn |
| `progression_transition` | bằng chứng hợp lệ dẫn tới trạng thái/hành động kế tiếp giải thích được |
| `variant_context` | nhiệm vụ/section/ngữ cảnh Academic/GT được biểu diễn và lấy mẫu |
| `delivery_mode_readiness` | có chuẩn bị/tương tác riêng cho hình thức thi được yêu cầu khi cần |
| `machine_contracts` | interface liên đơn vị chính xác, materialization định danh chung, tương thích lệch phiên bản, bảo toàn khả năng áp dụng, tuân thủ và kiểm chứng drift đều đạt |
| `rights_privacy_security` | yêu cầu áp dụng về nguồn/dữ liệu/đồng ý/truy cập/bí mật/truyền tải/lưu trữ/browser/SSE/service nội bộ/nhà cung cấp/biên công khai đều đạt |
| `reliability_recovery` | vòng đời, giao dịch/idempotency/concurrency, async bền vững, retry/backpressure, migration/deploy, khôi phục/vòng đời dữ liệu và suy giảm/khôi phục áp dụng đều đạt |
| `accessibility_capture_quality` | lỗi truy cập/thu dữ liệu/browser/thiết bị không bị hiểu nhầm là năng lực người học; tương tác được hỗ trợ phải dùng được |
| `cost_abuse_operations` | đạt giới hạn đã công bố về tần suất/lạm dụng/quota nhà cung cấp/compute/lưu trữ/traffic/capacity/scaling/sử dụng ngoài |
| `observability_audit` | đủ log/metric/correlation có cấu trúc, audit đặc quyền, provenance hệ quả, cảnh báo và tái dựng |
| `validation` | có bằng chứng kết quả thực nghiệm để nâng lên VALIDATED |

Thiếu artifact, chưa có triển khai/dữ liệu, sự bất tiện hoặc thiếu công nghệ có tên không bao giờ là lý do `NOT_APPLICABLE` hợp lệ. Chỉ bắt buộc content generator khi release trong phạm vi thực sự dựa vào việc tạo nội dung.

### Khả năng đi đến đích và khép kín dạng thi chính thức được suy ra

Độ bao phủ được suy ra từ ID/điều kiện áp dụng chính thức, không lấy từ ma trận phần trăm duy trì thủ công:

```text
Skill / enabling Knowledge where applicable
→ Curriculum route
→ trainable Practice/intervention
→ measurable Assessment path where required
→ Progression consequence
→ required product/content/runtime conditions
```

Với mỗi mục tiêu/đối tượng trong phạm vi: thiếu cạnh/đường đi áp dụng sẽ chặn; N/A tường minh do owner suy ra được phép; khả năng áp dụng chưa giải quyết vẫn là UNKNOWN. Không phải mọi đối tượng đều cần mọi cạnh, nhưng mục đích downstream bắt buộc phải đến được. Báo cáo suy ra là artifact kiểm chứng, không phải SSOT để chỉnh thủ công.

Độ bao phủ Skill và độ bao phủ dạng thi chính thức độc lập. Một Skill rộng hoặc một template điền chỗ trống không chứng minh bao phủ mọi dạng bên ngoài hay Presentation Class. Mục tiêu standard đầy đủ yêu cầu từng dạng `IELTS-*-*` áp dụng, Presentation Class bắt buộc, ngữ cảnh biến thể, tương tác, nguồn nội dung, đường chấm/evaluator và hệ quả bằng chứng được khép kín độc lập. Dạy tốt năng lực dùng chung không che được đường tương tác/nội dung dạng thi còn thiếu.

Trạng thái bằng chứng người học, lỗi thu dữ liệu, hiệu chuẩn evaluator, sự sẵn có nhà cung cấp, entitlement trả phí, ủy quyền đặc quyền và CoverageGap sản phẩm phải tách biệt. Bằng chứng người học thiếu/cũ/mâu thuẫn/thấp tự nó không phải CoverageGap; entitlement trả phí không thể biến lộ trình chưa bao phủ thành được hỗ trợ; thiếu tính năng trả phí tùy chọn không khiến lộ trình phổ thông được hỗ trợ trở thành không được hỗ trợ; thiếu năng lực sản phẩm phải là CoverageGap tường minh thay vì điểm yếu người học.

### Tuyên bố hỗ trợ mục tiêu và các cổng nâng trạng thái

Hỗ trợ tương đối theo mục tiêu áp dụng cho một tập ràng buộc mục tiêu, không phải một số Band. Cần biến thể chuẩn được giải quyết và đường được bao phủ cho mọi điều kiện TargetProfile thực sự bắt buộc: ràng buộc Band overall/từng kỹ năng, mục đích sử dụng bên ngoài, điều kiện đủ One Skill Retake đã chọn khi tuyên bố có bao gồm, và hình thức thi quan trọng. Nếu chỉ biết overall Band, kế hoạch có thể dùng hồ sơ làm việc được gắn nhãn không có thẩm quyền, nhưng không được tự tạo điểm tối thiểu từng kỹ năng thật. Xếp hạng không xóa điều kiện tiên quyết, tương thích dạng/ngữ cảnh/hình thức thi, điều kiện giao bài, tính đúng của bằng chứng hoặc CoverageGaps.

`TargetSupportDeclaration` ít nhất phải nêu:

- phạm vi mục tiêu/biến thể chính xác;
- các điều kiện Band/TargetProfile được hỗ trợ;
- hình thức thi và ràng buộc mục đích được hỗ trợ khi quan trọng;
- phiên bản sản phẩm/release;
- kết quả bao phủ dạng chính thức và định dạng phụ quan trọng;
- bao phủ feature/practice;
- phiên bản content manifest và phạm vi revision/danh mục đang hoạt động chính xác;
- tham chiếu chính sách/kết quả validation nội dung áp dụng;
- cổng tính mới/độc lập khi giao bài nếu quan trọng;
- cổng sự cố/khôi phục/ngừng dùng nội dung;
- phiên bản EvidenceRequirement/chính sách;
- trạng thái evaluator/hiệu chuẩn;
- phiên bản machine contract và tương thích theo chiều cho rollout đã chọn;
- trạng thái quyền riêng tư/quyền sử dụng/bảo mật;
- trạng thái kích hoạt bên thứ ba;
- trạng thái độ tin cậy/khôi phục/vòng đời dữ liệu;
- cổng khả năng tiếp cận/chất lượng thu dữ liệu;
- cổng chi phí/lạm dụng/vận hành;
- cổng quan sát/audit, gồm provenance phiên bản/cấu hình/nhà cung cấp có hệ quả và mục tiêu vận hành/chủ sự cố của release candidate;
- backlog validation đã biết nhưng không chặn;
- điều kiện thu hồi.

Hỗ trợ có phiên bản và có thể thu hồi khi cấu trúc thi, hình thức thi, nhà cung cấp, quyền/bảo mật, độ tin cậy/khôi phục, chi phí/năng lực tải, hiệu chuẩn, bao phủ/chất lượng/vận hành nội dung, quan sát, tương thích contract hoặc bằng chứng validation thay đổi đáng kể. `SUPPORTED_FOR_PRODUCT` nghĩa là có đường hoàn chỉnh đạt điều kiện release cho phạm vi công bố; không bao giờ bảo đảm kết quả IELTS bên ngoài hoặc cho phép tuyên bố kế hoạch bảo đảm Band mục tiêu hay một bài tập gây tăng Band khi thiếu cơ sở thực nghiệm phù hợp.
