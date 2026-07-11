# Constitution

## Core Laws

### Law 1: Implementation MUST conform to Authority
Mọi dòng code, mọi quyết định kiến trúc, mọi dependency đều PHẢI
truy nguyên được về một Authority File tương ứng. Nếu một chi tiết
triển khai không thể truy về Authority, nó không hợp lệ.

### Law 2: Missing Authority STOPS implementation
Nếu không có Authority cho một quyết định bắt buộc, việc triển khai
PHẢI dừng lại ngay lập tức. Khoảng trống đó PHẢI được ghi nhận vào
open-decisions.md và chờ Founder hoặc Vision Assistant phê duyệt.
Không được viết bất kỳ code nào để lách một Authority còn thiếu.

### Law 3: Implementation may NEVER redefine Authority
Các Authority Files (SSOT) không bao giờ được sửa đổi để hợp thức hóa
một triển khai đã được thực hiện. Triển khai phải tuân theo SSOT.
Nếu SSOT cần thay đổi, quyết định phải được phê duyệt TRƯỚC,
và chỉ sau đó Authority mới được cập nhật.

### Law 4: Sandbox Boundary
AI Runtime chỉ được hoạt động trong phạm vi workspace của repository.
Truy cập file bên ngoài workspace là FORBIDDEN.
Biến môi trường (API keys, tokens) KHÔNG được sao chép hoặc truyền ra ngoài.

### Law 5: Framework Immutability
Code trong node_modules/ là bất khả xâm phạm.
Code sinh bởi CLI tools (shadcn/ui, Drizzle Kit) chỉ được sửa qua CLI hoặc ADR.

## Gap Protocol

1. **DETECT** - Detect a gap: Implementation needs Authority but doesn't have one
2. **RECORD** - Record gap in open-decisions.md with classification
3. **STOP** - STOP implementation immediately
4. **RESUME** - Await Founder/Vision Assistant approval before resuming

## Loop Protocol

1. **IDENTIFY** - Identify a pattern or issue during implementation
2. **UPDATE** - Update Authority file to reflect findings
3. **VERIFY** - Verify update doesn't violate any Core Laws
4. **RECORD** - Record decision as ADR if it's architectural

## Gap Classification

| Category | P0 | P1 | P2 |
|----------|-----|-----|-----|
| **Product** | Core functionality missing | Non-core features | Nice-to-have |
| **Architecture** | Blocking architecture | Important but not blocking | Optimization |
| **Technical** | Breaking technical debt | Technical improvements | Minor refactor |

**P0 (Critical)**: STOP implementation immediately
**P1 (High)**: Pause until resolved
**P2 (Low)**: Document and defer
