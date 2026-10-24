# LEARNER-EXPERIENCE-CONTRACT Hợp đồng chất lượng sản phẩm dành cho người học

> **THẨM QUYỀN PRODUCT CHÍNH THỨC**
>
> Tài liệu này có thẩm quyền trong domain PRODUCT theo `docs/catalog/project.json`. Shard này chỉ quy định các ngữ nghĩa trải nghiệm bền vững hướng tới người học; không ấn định pixel, hệ thống thương hiệu cuối cùng, thư viện component hay bố cục phụ thuộc cách triển khai.

## Mô hình tư duy và hiển thị thông tin theo mức độ

Trải nghiệm mặc định là giao diện đơn giản đặt trên hệ thống có chiều sâu. Giao diện thể hiện mục tiêu do người học khai báo, trạng thái hiện tại dựa trên bằng chứng, một số ít trở ngại quan trọng hoặc điều chưa biết, cùng hành động hợp lệ tiếp theo; người học không phải hiểu ontology, nhà cung cấp, schema hay cơ chế chính sách nội bộ.

Điều hướng chính tập trung vào **Today, Skills, Practice, Review, Media, Progress và Mock**; Target/Diagnostic, Vocabulary và Profile được hiển thị khi ngữ cảnh cần. Màn hình mặc định phải trả lời được điều gì quan trọng lúc này. Có thể mở dần bản đồ kỹ năng/khoảng trống, Knowledge Objects, lịch sử lần làm bài/bằng chứng, chính sách, nguồn gốc hoặc giải thích sâu hơn khi người học yêu cầu.

ID nội bộ, cơ chế đồ thị, kết nối evaluator/nhà cung cấp, trạng thái triển khai và chi tiết migration/catalog không thuộc giao diện thông thường. Không được tự tạo tỷ lệ hoàn thành tổng thể khi bằng chứng quan trọng, phạm vi hỗ trợ sản phẩm hoặc điều kiện áp dụng vẫn chưa được xác định.

## Đặc trưng thị giác và ngôn ngữ dành cho người học

Phong cách bền vững lấy kiểu chữ làm chủ đạo, ưu tiên nội dung, điềm tĩnh và dễ đọc. Phân cấp, khoảng cách, chữ rõ ràng, trạng thái minh bạch và hành động học tập quan trọng hơn mật độ trang trí.

Tránh các mẫu thiết kế áp dụng toàn sản phẩm làm lu mờ nhiệm vụ học hoặc biến hệ thống thành màn trình diễn: phong cách edtech trẻ con, game hóa trang trí, gradient/kính mờ quá mức, hiệu ứng AI phát sáng tương lai, các ô dashboard chằng chịt, huy hiệu/icon rối mắt, điều hướng theo ontology và giao diện gắn thương hiệu AI chung chung.

Nội dung giao diện phải nêu **năng lực, hành động, kết quả, trạng thái, trở ngại, điều chưa chắc chắn hoặc cách khôi phục trước tiên**. Thương hiệu AI hay nhà cung cấp không thay thế được nhãn sản phẩm có ý nghĩa. Vẫn phải công bố thông tin về quyền riêng tư, nguồn gốc, nguồn nội dung, evaluator/AI hoặc yếu tố khác khi pháp luật hay hậu quả thực tế yêu cầu; ưu tiên ngôn ngữ về năng lực không đồng nghĩa che giấu nguồn gốc quan trọng.

## Biểu tượng, hình ảnh và chiều sâu

Icon phải có chức năng rõ ràng, tiết chế, nhất quán, đi kèm chữ hoặc phương án hỗ trợ tiếp cận khi ý nghĩa có thể mơ hồ. Không dùng emoji, icon hoạt hình, bộ biểu tượng trộn lẫn hay đặt icon khắp nơi chỉ để trang trí.

Hình ảnh được dùng ít nhưng chất lượng cao. Chiều sâu nền có thể hỗ trợ phân cấp, nhưng nội dung phải ở lớp trước và không thông tin bắt buộc nào chỉ được truyền tải qua nền. Loại trừ hình công nghệ sáo rỗng, robot tương lai và ảnh hero cạnh tranh sự chú ý với nhiệm vụ học tập.

## Chuyển động

Chỉ sử dụng chuyển động khi làm rõ thay đổi trạng thái, phân cấp, phản hồi hoặc điều hướng; không dùng làm màn trình diễn. Tránh parallax, chuyển động trang trí liên tục, hiệu ứng dài cản trở thao tác và chuyển động gây mất ổn định khi đọc hay điều khiển.

Ưu tiên giảm chuyển động, khả năng tiếp cận, giao diện thích ứng, bố cục ổn định và sự liên tục của nhiệm vụ hơn hiệu ứng. Chế độ không chuyển động hoặc giảm chuyển động phải giữ nguyên ý nghĩa sản phẩm và khả năng thực hiện hành động.

## Chất lượng tương tác ở mức chi tiết

Những trạng thái hiển thị cho người học phải được phân biệt rõ và sử dụng được:

- trạng thái đang tải và skeleton giữ ổn định bố cục, không giả vờ rằng nội dung chưa biết đã được xác định;
- trạng thái trống cho biết thứ gì đang thiếu và hành động hợp lệ tiếp theo nếu có;
- lỗi, suy giảm chất lượng, bị vô hiệu hóa, đang chờ và không khả dụng là các trạng thái khác nhau, không gom thành một lỗi chung;
- hover, focus, nhấn, chọn, mở rộng và vô hiệu hóa phải dễ nhận biết, nhất quán với phương thức nhập liệu hỗ trợ chúng;
- thao tác bàn phím, thứ tự/khả năng nhìn thấy focus và ngữ nghĩa trình đọc màn hình phải hợp lệ; chỉ màu sắc hoặc icon không đủ truyền đạt khác biệt bắt buộc;
- câu chữ ngắn gọn vẫn phải thể hiện sự chưa chắc chắn, không dựng lên sự chắc chắn chỉ để giao diện gọn;
- hành động chính, phụ và khôi phục phải phân biệt được; và
- bố cục responsive/mobile phải giữ phân cấp, khả năng đọc, điều khiển ổn định và thu/nhập liệu an toàn, thay vì chỉ thu nhỏ dashboard máy tính.

Chất lượng tương tác chi tiết là một phần của tính đúng đắn sản phẩm khi trạng thái thiếu hoặc mơ hồ có thể khiến người học hiểu sai bằng chứng, mức hỗ trợ, tính đủ điều kiện, việc thu dữ liệu thành công hay hành động hợp lệ kế tiếp.

## Bất biến: giao diện đơn giản, hệ thống chuyên sâu

Nội bộ có thể duy trì cơ chế chuyên sâu về bằng chứng, điều kiện tiên quyết, độ bao phủ, nội dung, nhà cung cấp, vòng đời và nguồn gốc; nhưng giao diện mặc định vẫn tập trung vào những điểm ngắn gọn và quan trọng: mục tiêu, trạng thái hiện tại, trở ngại/điều chưa biết quan trọng và hành động hợp lệ kế tiếp.

Thông tin giải thích sâu được mở dần khi giúp người học hiểu vì sao có đề xuất, bằng chứng nào hỗ trợ trạng thái, điều gì chưa được giải quyết hoặc cách khôi phục. Việc mở dần thông tin không được che giấu trở ngại làm thay đổi hành động hợp lệ; sự đơn giản không được gộp những trạng thái khác nhau về bản chất thành một điểm số thân thiện.

Hợp đồng này quản lý ngữ nghĩa chất lượng sản phẩm mang tính bền vững. Cỡ chữ, màu sắc, bo góc, thang khoảng cách, thời gian chuyển động, breakpoint, triển khai component, đồ họa cuối cùng và token thương hiệu là lựa chọn của implementation/design system, miễn giữ nguyên các ngữ nghĩa này.
