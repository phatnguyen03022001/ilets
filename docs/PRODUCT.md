# Product

> **THẨM QUYỀN PRODUCT CHÍNH THỨC**
>
> Có thẩm quyền trong domain PRODUCT theo `docs/catalog/project.json`. `CONSTITUTION.md` và `OBJECTIVE.md` giữ thẩm quyền riêng; `contracts/**` giữ thẩm quyền machine contract chính xác trong phạm vi của nó. Những tham chiếu tới `spec/**` hoặc `design/**` bên dưới, kể cả cách gọi “canonical” trong lịch sử, chỉ để truy nguyên nguồn gốc, không tạo thẩm quyền ngang hàng.

Tài liệu này lưu giữ sự thật sản phẩm IELTS đã được chấp nhận trong domain `PRODUCT` của mô hình `agent-documents` V1. Các công việc lịch sử `TASK-0012`, `TASK-0013` và `TASK-0025` đã xác lập actor/RBAC/entitlement, bộ tính năng đầy đủ và ranh giới tham gia của người giám hộ/người đại diện pháp lý/người trả tiền hiện được lưu ở đây. `docs/catalog/project.json` sở hữu ID catalog, quan hệ có kiểu, coverage và trạng thái phạm vi milestone hiện tại. File này không sở hữu độc lập trạng thái có thể thay đổi đó, không kích hoạt nhà cung cấp và không thay đổi ngữ nghĩa sản phẩm đã khóa.

## Objective

Xây dựng đặc tả hệ thống học IELTS đầy đủ, dựa trên bằng chứng và thiết kế sản phẩm/runtime cho người tự học có mục tiêu IELTS cụ thể. Sản phẩm điều phối lộ trình từ mục tiêu, qua bằng chứng hiện tại và các khoảng trống/điều chưa biết được nêu rõ, đến hành động học hợp lệ kế tiếp, bằng chứng mới trung thực và lập lại kế hoạch tự động.

Giá trị bền vững nằm ở vòng lặp được quản trị từ mục tiêu → bằng chứng → hành động tiếp theo. Từng mô hình AI, thuật toán SRS, nhà cung cấp media, evaluator, định dạng bài tập và cơ chế tương tác thời gian thực là năng lực có thể thay thế, không phải trung tâm ngữ nghĩa sản phẩm.

Kết quả mong muốn là lộ trình được quản trị từ đầu vào có cấu trúc ở Band 3 đến Band 9 trong Listening, Reading, Writing và Speaking, giữ khả năng lệch trình độ giữa kỹ năng, trạng thái bằng chứng chưa biết/cũ/mâu thuẫn tường minh, ràng buộc overall/từng kỹ năng, khác biệt Academic và General Training, và phân biệt trung thực tiến bộ học, readiness, hỗ trợ sản phẩm với điểm IELTS bên ngoài thực tế.

Sản phẩm có thể cung cấp lộ trình quản trị mạnh hướng tới mục tiêu, nhưng không được ngụ ý rằng sử dụng hệ thống sẽ đảm bảo một điểm IELTS bên ngoài trong tương lai.

**Nguồn:** `OBJECTIVE.md` — Mục đích, kết quả mục tiêu; `spec/00-PRODUCT.md` — Định danh sản phẩm, ranh giới lời hứa.

## Actors and roles

Sản phẩm phân biệt ba khái niệm:

1. **người tham gia / lớp actor**: ai hoặc hệ thống nào tham gia việc học, hỗ trợ, nội dung hoặc vận hành;
2. **authorization**: dựa trên capability, các bundle vai trò có tên chỉ làm tập capability mặc định thuận tiện cho người vận hành;
3. **commercial entitlement**: kiểm soát khả năng truy cập trải nghiệm sản phẩm đủ điều kiện và không bao giờ là vai trò RBAC.

`LEARNER`, `COLLABORATOR`, `REVIEWER`, `ADMIN` và `OWNER` là bundle ủy quyền mặc định thông thường. Vai trò AI support và chuyên gia tùy chọn đang có không tự thân là bundle ủy quyền. Định danh service/system hoặc con người chỉ được nhận bundle hoặc tập capability hẹp hơn khi được cấp tường minh và có thẩm quyền; tham gia không mặc nhiên tạo đặc quyền.

**Nguồn:** `design/04-application-flows.md` — Actor, capability và mô hình RBAC; `design/00-learning-experience.md` — Ranh giới trải nghiệm an toàn cho người chưa thành niên, khả năng sử dụng theo entitlement.

### Self-directed IELTS learner

Người học chính có mục tiêu IELTS thực tế và sử dụng lộ trình sản phẩm bình thường, không bắt buộc giáo viên. Người học khai báo mục tiêu và có thể đổi sau đó. Điều khiển Swap, Skip, Shorten hoặc Change-skill được chọn trong những lộ trình hợp lệ nhưng không được xóa điều kiện Required, điều kiện mục tiêu chưa được hỗ trợ hoặc khoảng trống bao phủ sản phẩm. Bundle mặc định `LEARNER` chỉ cấp hoạt động người học/tài khoản thông thường trong phạm vi của họ, không mặc nhiên cấp capability đặc quyền về nội dung, hỗ trợ, entitlement, chính sách, bảo mật hay quản trị vai trò.

**Nguồn:** `OBJECTIVE.md` — Mục đích, lộ trình người học; `spec/00-PRODUCT.md` — Định danh sản phẩm, ranh giới hỗ trợ con người.

### Guardian or legal representative

Khi áp dụng, người giám hộ hoặc người đại diện pháp lý là actor con người tách biệt với người học và người thanh toán. Quan hệ hoặc tư cách đại diện tự nó không tạo bundle vai trò RBAC, quyền capability, entitlement thương mại hay thẩm quyền học tập/bằng chứng.

Chỉ có tư cách giám hộ/đại diện không cấp truy cập không giới hạn vào lịch sử người học, bản ghi âm, hội thoại AI, bằng chứng, thông tin khoảng trống/điểm yếu hoặc dữ liệu bảo vệ khác. Mọi truy cập dữ liệu bảo vệ và hành động đồng ý, xuất, xóa nhân danh người khác phải giới hạn theo mục đích, chính sách và các quyền thực sự áp dụng cho hành động đó. Minh bạch với người học và giảm tối đa tiết lộ không cần thiết vẫn là ràng buộc sản phẩm.

Ngưỡng tuổi, tư cách chưa thành niên, điều kiện làm đại diện, yêu cầu đồng ý và quyền đại diện thay đổi theo khu vực pháp lý; chúng là bằng chứng chính sách/tuân thủ bên ngoài có thể thay đổi, không phải ngữ nghĩa IELTS đã đóng băng. Tài liệu không đưa ra kết luận pháp lý.

**Nguồn:** `design/00-learning-experience.md` — Ranh giới trải nghiệm an toàn cho người chưa thành niên.

### Payer

Khi áp dụng, người thanh toán là actor thương mại con người, tách khỏi người học và người giám hộ/đại diện pháp lý. Trả tiền hoặc tài trợ truy cập không tự tạo định danh người học, bundle RBAC, quyền capability, thẩm quyền học/bằng chứng hoặc quyền xem dữ liệu người học được bảo vệ.

Tư cách người trả tiền không cấp truy cập không giới hạn vào lịch sử, bản ghi âm, hội thoại AI, bằng chứng, khoảng trống/điểm yếu hoặc thông tin bảo vệ của người học. Entitlement thương mại chỉ được thay đổi khả năng sử dụng sản phẩm hợp lệ sau đối soát có thẩm quyền; nó không biến người trả tiền thành người có quyền truy cập dữ liệu người học hay đặc quyền quản trị.

**Nguồn:** `design/00-learning-experience.md` — An toàn cho người chưa thành niên, khả năng truy cập theo entitlement; `spec/00-PRODUCT.md` — Ranh giới truy cập thương mại.

### AI support role

AI có thể hỗ trợ giải thích, mô hình hóa, luyện có hướng dẫn, phản hồi, kèm học, tạo, phân tích, đánh giá, tương tác Speaking và học với media trong vai trò giới hạn. AI không trở thành sự thật IELTS, sự thật Assessment, thẩm quyền Progression hoặc chuẩn định nghĩa năng lực người học.

Đây là vai trò tham gia, không tự động là bundle ủy quyền. Định danh AI/service do hệ thống quản lý chỉ được nhận `COLLABORATOR`, `REVIEWER` hoặc quyền capability hẹp hơn tường minh khi chính sách ủy quyền hiện hành cho phép; việc hỗ trợ bằng AI không hàm ý quyền release, entitlement, dữ liệu bảo vệ, quản trị cấp quyền, bảo mật hay đặc quyền khác.

**Nguồn:** `OBJECTIVE.md` — Trải nghiệm sản phẩm; `spec/00-PRODUCT.md` — AI hỗ trợ mà không thành tiêu chuẩn, ranh giới hỗ trợ con người; `design/04-application-flows.md` — Actor, capability, RBAC.

### Optional human expert role

Đầu vào chuyên gia con người có thể dùng cho sở thích hoặc huấn luyện, hoặc trở thành bằng chứng bên ngoài khi quy tắc Assessment cho phép. Không bắt buộc con người chấm điểm để có lộ trình học thông thường được hỗ trợ chỉ nhằm dựng lên hệ quả sản phẩm không thể hỗ trợ an toàn.

Hỗ trợ học tập từ chuyên gia tùy chọn cũng là vai trò tham gia, không tự động là bundle đặc quyền.

**Nguồn:** `spec/00-PRODUCT.md` — Ranh giới hỗ trợ con người.

### Privileged human operator

Người vận hành đặc quyền là người đã xác thực, hành động qua một hoặc nhiều capability tường minh về nội dung, vận hành, hỗ trợ, entitlement, chính sách, bảo mật hoặc quản trị cấp quyền. Đặc quyền chỉ được thực thi qua ranh giới ủy quyền do Core sở hữu; không trở thành thẩm quyền về học tập, Band, Assessment, bằng chứng, chất lượng nội dung hay bỏ qua validation.

**Nguồn:** `design/04-application-flows.md` — Actor, capability, RBAC.

### Authorization bundles

Sự thật về quyền là tập capability có hiệu lực; tên dưới đây là bundle mặc định ổn định, không phải tầng thẩm quyền thứ hai:

| Bundle | Mục đích mặc định | Ranh giới |
| --- | --- | --- |
| `LEARNER` | hoạt động sản phẩm thông thường của người học/tài khoản đã xác thực | không quản trị đặc quyền nội dung, hỗ trợ, entitlement, chính sách, bảo mật hay cấp quyền |
| `COLLABORATOR` | soạn/tạo/sửa ứng viên nội dung; tạo/xử lý có giới hạn; phân loại, khử trùng lặp, đề xuất sửa và kiểm tra không hạn chế theo yêu cầu | không release/kích hoạt, thay đổi entitlement, truy cập dữ liệu người học được bảo vệ, quản trị cấp quyền hoặc vận hành nhạy cảm bảo mật |
| `REVIEWER` | quyền `COLLABORATOR` cộng kiểm tra/review nội dung, cung cấp đầu vào validation, chấp thuận/từ chối được phép, yêu cầu tái validation và cách ly giao bài mới khi có sự cố | không vượt qua cổng cứng, cấp vai trò hoặc quản trị người học/bảo mật không liên quan |
| `ADMIN` | release/retire nội dung có cổng và xử lý sự cố; hỗ trợ người dùng giới hạn; đối soát entitlement; quan sát vận hành; quản trị chính sách vận hành đã duyệt; truy cập dữ liệu người học bảo vệ theo mục đích chỉ khi được ủy quyền riêng | không đổi sự thật học/bằng chứng, override validation chung, tăng trần tài chính tối quan trọng hoặc sở hữu đặc quyền/chính sách bảo mật cấp cao nhất |
| `OWNER` | quyền `ADMIN` cộng cấp/thu hồi vai trò/capability và vận hành bảo mật nhạy cảm, khôi phục, phá hủy hoặc trần tài chính tối quan trọng khi được cấp tường minh | vẫn không thể viết lại lịch sử bất biến hoặc vượt cổng cứng về nội dung/bằng chứng/sự thật sản phẩm |

Một actor có thể giữ nhiều quyền tương thích; ưu tiên quyền capability hẹp tường minh nếu bundle quá rộng. Phê duyệt của `REVIEWER` là đầu vào validation vận hành, không phải thẩm quyền Assessment; release chỉ hợp pháp sau khi các cổng áp dụng đều đạt.

**Nguồn:** `design/04-application-flows.md` — Actor, capability, RBAC.

### Commercial entitlement

Entitlement thương mại là chiều độc lập của khả năng dùng sản phẩm. Authority này có thể mô tả trạng thái hiển thị `FREE`, `PRO` nhưng chúng không phải ID `ROL-*`, và V1 không định nghĩa lớp định danh entitlement. Entitlement có thể mở khóa trải nghiệm chi phí cao vốn đã hợp lệ, nhưng không được thay đổi sự thật IELTS/Band/bằng chứng, ngữ nghĩa mục tiêu/tiên quyết, lịch sử/quyền dữ liệu người học hay đặc quyền quản trị.

Lộ trình không cần trả phí vẫn phải là hệ thống học IELTS sử dụng được thực sự. Entitlement `PRO` hợp lệ chỉ được mở realtime AI Speaking sau đối soát entitlement có thẩm quyền và khi toàn bộ cổng nhà cung cấp/capability/hỗ trợ sản phẩm thực sự đạt. Giá, quota, hành vi nhà cung cấp thanh toán và đóng gói thương mại cụ thể nằm ngoài authority PRODUCT này.

**Nguồn:** `design/00-learning-experience.md` — Truy cập theo entitlement, trải nghiệm Speaking; `design/04-application-flows.md` — vòng đời entitlement người học so với ủy quyền vận hành.

## Product-level features and capabilities

Bề mặt tính năng đầy đủ có ý nghĩa đã được định rõ. `FTR-001..FTR-006` là anchor macro lộ trình lịch sử cho traceability đã chấp nhận; không tính là đơn vị tính năng chi tiết. Danh mục chi tiết:

| Dải FTR chi tiết | Alias nguồn lịch sử | Số năng lực đáng kể |
| --- | --- | ---: |
| `FTR-007..FTR-014` | `L-F01..L-F08` Listening | 8 |
| `FTR-015..FTR-022` | `R-F01..R-F08` Reading | 8 |
| `FTR-023..FTR-031` | `W-F01..W-F09` Writing | 9 |
| `FTR-032..FTR-040` | `S-F01..S-F09` Speaking | 9 |
| `FTR-041..FTR-046` | `X-F01..X-F06` dùng chung | 6 |
| `FTR-047..FTR-053` | nhóm giao diện đặc quyền/admin/BOPS | 7 |

40 ID dành cho người học là bề mặt năng lực được đặt tên chính xác, kế thừa từ `design/01-skill-features.md`. Bảy ID đặc quyền là các nhóm giao diện quan trọng tối thiểu được `TASK-0012` xác lập và giữ từ `design/04-application-flows.md`: nội dung/review/release; hỗ trợ người dùng; đối soát entitlement; quan sát vận hành/nhà cung cấp/công việc; quản trị chính sách đã duyệt; quản trị bảo mật/truy cập; và audit.

Cách phân rã này cố ý **không** làm phình danh mục FTR bằng các khái niệm lân cận:

- **Today, Skills, Practice, Review, Media, Progress, Mock** là điều hướng/kiến trúc thông tin, không phải bảy ID tính năng thêm;
- 28 mode `PM-*` là cách đóng gói thực hành theo `design/02-practice-catalog.md`, không phải 28 FTR;
- Skill Leaves / micro-skills cho người học và Knowledge Objects là ontology học tập, không phải ID tính năng;
- nguồn media audio/image, trạng thái transcript/quyền và nhà cung cấp nguồn thuộc trạng thái source/dependency theo `docs/behavior/practice-and-media.md`, `docs/DATA.md` và `docs/interfaces/provider-lifecycle-and-routes.md`;
- `FREE`/`PRO` là trạng thái entitlement thương mại, không phải tính năng/vai trò;
- từng nhà cung cấp AI/LLM/STT/TTS hoặc đường model là dependency/capability có thể thay, không phải tính năng;
- kèm học bằng AI là lớp hỗ trợ/hình thức cung cấp xuyên tính năng; `FTR-006` vẫn là anchor macro hỗ trợ giới hạn, không phải tính năng chi tiết thứ 41 bị nhân đôi.

Khác biệt Academic/General Training và hình thức thi chỉ cấu hình cùng tính năng chi tiết khi tương tác/năng lực người học về bản chất dùng chung. Chúng không tạo ID tính năng trùng hoặc chuẩn Band khác. One Skill Retake tập trung một kỹ năng đã chọn, không là kỹ năng thứ năm hay nhóm tính năng nhân đôi.

Tên chi tiết, alias nguồn, nhóm flow và nhóm acceptance thuộc `docs/BEHAVIOR.md` cộng `docs/catalog/project.json`. Quan hệ DATA/IFC/EXT/CAP chi tiết đã được giải quyết trong các migration theo domain giới hạn ở `TASK-0014..TASK-0017`; ánh xạ giữ tham chiếu đã chấp nhận chính xác hoặc N/A được chứng minh nơi phù hợp, không thêm quan hệ domain lân cận mang tính phỏng đoán.

**Nguồn:** `design/01-skill-features.md`; `design/02-practice-catalog.md` (danh mục practice lịch sử); `docs/behavior/practice-and-media.md` (media eligibility hiện tại); `design/04-application-flows.md`.

## Scope

**Media eligibility:** audio và image là hai kiểu media duy nhất được hỗ trợ cho bytes hệ thống quản lý và tham chiếu bên ngoài. Không hỗ trợ upload, lưu trữ, nhập video hay YouTube, liên kết dùng làm video, embed/player/stream, extraction, transcript-from-video hoặc provider route. Văn bản Reading, đầu vào Writing, transcript/caption/đáp án audio được phép và dữ liệu chart/diagram có cấu trúc vẫn là nội dung hợp lệ không phải video, không buộc lưu như binary audio/image. Học qua media tiếp tục chịu các cổng về quyền, nguồn gốc và bằng chứng.

Trạng thái phạm vi milestone tài liệu có thể thay đổi hiện thuộc riêng `docs/catalog/project.json`; Markdown này cố ý không lặp `OPEN`, `FROZEN` hoặc `SCOPE_OPEN`. Công việc quan hệ DATA/IFC/EXT/CAP có giới hạn và `UNK-001..UNK-009` đã giải quyết. `DOCS_READY` được mô hình tài liệu đã pin suy ra từ toàn bộ bộ chính thức cùng trạng thái catalog; không thuộc sở hữu hoặc lưu trong file này.

Phạm vi học IELTS chuẩn đầy đủ dự định:

```text
IELTS Academic
+ IELTS General Training
+ Listening / Reading / Writing / Speaking
+ official task/question families
+ Band-3→9 structured learning paths
+ target profile → diagnostic → learning → practice → assessment → readiness
```

Đặc tả học tập hiện hành bao gồm bốn kỹ năng chấm điểm; grammar, vocabulary và phonology nền; Skill Leaves/Knowledge Objects ổn định; kỳ vọng Band/tiêu chí đầu ra; thứ tự và điều kiện tiên quyết chương trình học; learning mechanisms và Practice Types dùng lại; ngữ nghĩa Assessment, bằng chứng, readiness, mastery, progression, regression, review, certification; và biểu diễn nội dung ở mức khái niệm cần để hiện thực hệ thống.

Academic có thể release trước General Training nhưng thứ tự release không định nghĩa lại cấu trúc học đầy đủ mong muốn. Listening và Speaking dùng chung. Năng lực và chiến lược Reading chung, còn ngữ cảnh/cách chấm theo biến thể là điều kiện bên ngoài/ngữ cảnh. Academic Writing Task 1 và General Training Writing Task 1 giữ khác biệt thật của bài thi; Writing Task 2 và tiêu chí Writing/grammar/lexis/cohesion chung không bị nhân đôi chỉ vì biến thể.

IELTS for UKVI Academic/General Training dùng lại cấu trúc học tương ứng, thêm điều kiện hành chính/bảo mật ngoài. One Skill Retake dùng lại cấu trúc kỹ năng được chọn. IELTS Life Skills là cấu trúc pass/fail chỉ Listening/Speaking khác biệt, không thuộc mục tiêu học sản phẩm hiện tại trừ khi được thêm tường minh sau này.

**Nguồn:** `OBJECTIVE.md` — Phạm vi IELTS chuẩn đầy đủ, Scope; `spec/00-PRODUCT.md` — Phạm vi học IELTS chuẩn.

## Non-goals

Authority này giữ các loại trừ hiện tại. Repo chưa đóng băng:

- nhà cung cấp cloud/hosting cuối cùng;
- nhà cung cấp PostgreSQL cuối cùng;
- nhà cung cấp object storage cuối cùng;
- nhà cung cấp identity cuối cùng;
- nhà cung cấp thanh toán cuối cùng;
- nhà cung cấp AI/LLM/STT/TTS cuối cùng;
- các bảng database production cuối cùng;
- Kubernetes hoặc kiến trúc nhiều vùng khi chưa chứng minh nhu cầu;
- thiết kế thị giác pixel-perfect hoặc hệ thống thương hiệu cuối cùng;
- hướng dẫn coding agent riêng theo nhà cung cấp.

Owner sản phẩm cũng không định nghĩa giá cụ thể, cơ chế thu hút người dùng hoặc lưu giữ, flow UI cụ thể, lựa chọn provider hoặc chi tiết triển khai kỹ thuật. Những quyết định đó có thể được chọn sau bên trong ranh giới provider/runtime chính thức nhưng không được định nghĩa lại sự thật học tập hoặc ngữ nghĩa hỗ trợ sản phẩm.

**Nguồn:** `OBJECTIVE.md` — Non-goals; `spec/00-PRODUCT.md` — Non-goals, ranh giới implementation, ranh giới truy cập thương mại.

## Domain and external constraints

- Yêu cầu IELTS bên ngoài là sự thật domain bên ngoài. Ontology học phải ánh xạ chúng sang năng lực mà không sao chép giao diện thi vào ontology hoặc để lựa chọn triển khai định nghĩa lại năng lực IELTS.
- Academic và General Training dùng chung cấu trúc, chỉ khác nơi IELTS thực sự khác. Hình thức thi có thể thay đổi sự quen thuộc giao diện, thao tác, luyện thời gian/điều hướng, điều kiện khả năng tiếp cận/thu dữ liệu hoặc điều kiện chấp nhận/đủ điều kiện bên ngoài, nhưng không định nghĩa lại Skill, yêu cầu Knowledge hoặc chuẩn chất lượng Band, trừ khi tuyên bố cụ thể về readiness cho điều kiện thi đó.
- Hệ thống học độc lập L1. Ngôn ngữ thứ nhất có thể giúp bản địa hóa, giải thích, ví dụ, giả thuyết can nhiễu và khắc phục, nhưng không tạo định nghĩa năng lực IELTS khác.
- Bằng chứng phải trung thực. Attempt, Observation, EvidenceFact, mastery/readiness, gaps và hành động kế tiếp tách biệt; thời gian học, bài hoàn thành, nội dung tiêu thụ, đầu ra AI hoặc điểm công cụ không tự chứng nhận năng lực.
- Mục tiêu người học khai báo được giữ cho tới khi họ đổi. Quyền chủ động cho phép chọn trong lộ trình hợp lệ nhưng không vượt Required hoặc điều kiện mục tiêu chưa được hỗ trợ/chưa bao phủ.
- Entitlement thương mại có thể mở truy cập capability chi phí cao nhưng không mua chuẩn Band khác, tính hợp lệ EvidenceFact mạnh hơn, bypass điều kiện tiên quyết ẩn, sự chắc chắn giả, đặc quyền quản trị hoặc quyền sở hữu lịch sử/quyền dữ liệu người học. Lộ trình không có cổng trả phí phải là hệ thống học IELTS thực sự dùng được.
- Nhà cung cấp media là nền nguồn, không là sự thật học. Phải tôn trọng transcript/quyền và xử lý lỗi an toàn; không mặc nhiên cho phép extraction/download tùy ý.
- Provider ngoài có thể thay thế sau ranh giới capability rõ ràng, không thành owner cạnh tranh của sự thật học tập.
- Coverage sản phẩm và hỗ trợ tách khỏi tính nhất quán mô hình học. Repo phân biệt `MODELLED → COVERED → SUPPORTED_FOR_PRODUCT → VALIDATED`; không được dùng tỷ lệ “100% IELTS” tổng thể che giấu điều kiện bắt buộc còn thiếu.
- Sản phẩm có thể kiểm soát thuộc tính quy trình như toàn vẹn mục tiêu, quy tắc tiên quyết/bằng chứng, ngữ nghĩa hành động kế tiếp giải thích được và trạng thái hiện tại trung thực. Không thể đảm bảo kết quả IELTS bên ngoài tương lai.

**Nguồn:** `OBJECTIVE.md` — Mục đích, lộ trình người học, coverage/hỗ trợ, media, thiết kế app/runtime; `spec/00-PRODUCT.md` — Nguyên tắc sản phẩm, ranh giới hình thức thi, ngôn ngữ/bản địa hóa, hỗ trợ con người, truy cập thương mại, implementation, tính hoàn chỉnh mô hình so với hỗ trợ, lời hứa.

## Ranh giới thẩm quyền và kết thúc

Không tìm thấy xung đột thẩm quyền trong các nguồn gốc PRODUCT được giữ lại cho lần chuyển đổi này.

Thẩm quyền tài liệu dự án mang tính quy phạm hiện tại là tám domain root `docs/*.md`, shard hợp lệ một cấp được catalog tham chiếu và `docs/catalog/project.json`, theo thứ tự đọc ưu tiên trong `README.md`. `CONSTITUTION.md` và `OBJECTIVE.md` giữ thẩm quyền riêng. `contracts/**` chỉ có thẩm quyền chính xác với machine contract trong scope. `spec/**`, `design/**` được giữ theo dõi làm nguồn gốc/tài liệu lịch sử, không có thẩm quyền ngang hàng kể cả khi ngôn ngữ trạng thái/owner lịch sử nói khác.

`docs/catalog/project.json` là owner duy nhất trạng thái phạm vi milestone tài liệu hiện tại; `UNK-001..UNK-009` và công việc quan hệ chi tiết DATA/IFC/EXT/CAP đã được giải quyết tại đó. `DOCS_READY`, khi mô hình đã pin suy ra, chỉ là bằng chứng đóng tài liệu. Kết thúc/chuyển đổi tài liệu không tự tạo DESIGN LOCK, trạng thái standards/assurance, kích hoạt nhà cung cấp, mức sẵn sàng implementation, promotion hay release.

**Nguồn:** `.agent/tasks/TASK-0013/task.yaml@07f409f37243f80695fcb98333f3ddd728960e7b`; `.agent/tasks/TASK-0025/task.yaml@0f07168ce65424b3411ac07d6b1fb122a97d2037`; `agent-documents@acb3f02616e700190586681306a86905792e4c07` — ngữ nghĩa feature/relation/unknown V1.

## Tóm tắt truy nguyên

| Vấn đề PRODUCT | Nguồn gốc lịch sử | Mục chính thức |
| --- | --- | --- |
| mục tiêu | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Objective |
| người tham gia / vai trò | `OBJECTIVE.md`; `spec/00-PRODUCT.md`; `design/00-learning-experience.md`; `design/04-application-flows.md` | Actors and roles |
| tính năng / năng lực | `design/01-skill-features.md`; `design/04-application-flows.md` | Product-level features and capabilities |
| phạm vi / loại trừ | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Scope; Non-goals |
| ràng buộc domain / ngoài | `OBJECTIVE.md`; `spec/00-PRODUCT.md` | Domain and external constraints |

Các tham chiếu migration PRODUCT gốc còn truy về snapshot nguồn đã ghi. Nền actor/RBAC/entitlement của `TASK-0012` kế thừa từ `design/00-learning-experience.md` và `design/04-application-flows.md`; ID tính năng chi tiết `TASK-0013` bắt nguồn từ `design/01-skill-features.md` và ranh giới giao diện/capability đặc quyền trong `design/04-application-flows.md` tại authority đích được chấp nhận `b60034a50d9a5ee5f197887ed14e7b917e919660`. `TASK-0025` khôi phục phân biệt giám hộ/đại diện pháp lý/người trả tiền an toàn cho người chưa thành niên từ `design/00-learning-experience.md` mà không tạo bundle RBAC giám hộ/người trả tiền hay subsystem triển khai. Các tham chiếu này chỉ là provenance; thẩm quyền hiện hành theo thứ tự docs/catalog/README bên trên.
