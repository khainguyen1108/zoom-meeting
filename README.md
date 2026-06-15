Voice-to-Voice hoạt động như thế nào?
Có 4 bước chính nối tiếp nhau:
Bước 1 — Thu âm & gửi lên server

Micro của người nói được capture qua WebRTC, cắt thành chunks nhỏ ~200ms và stream liên tục lên server. Không chờ hết câu mới gửi — gửi cuốn chiếu để giảm độ trễ.
Bước 2 — Speech-to-Text (STT)

Server dùng Whisper (OpenAI) hoặc Google Speech-to-Text nhận diện giọng nói thành văn bản. Dùng chế độ streaming — trả về partial results (kết quả tạm) ngay khi có vài từ, không chờ hết câu. Ví dụ: người nói "Xin chào mọi người hôm nay..." → STT trả về từng cụm từ liên tục.
Bước 3 — Dịch văn bản

Văn bản từ STT được gửi đến DeepL hoặc Google Translate. Cũng dịch theo streaming — dịch từng đoạn ngắn thay vì chờ cả câu hoàn chỉnh. Đánh đổi: đôi khi dịch hơi cụt vì chưa đủ ngữ cảnh, nhưng bù lại nhanh hơn.
Bước 4 — Text-to-Speech (TTS) → phát cho người nghe

Văn bản đã dịch được chuyển thành giọng nói AI bằng Google Neural TTS hoặc ElevenLabs. Audio được stream xuống thiết bị người nghe và phát ra loa — người nghe chọn ngôn ngữ mình muốn nghe trong settings.

Vấn đề khó nhất: Audio Mixing
Người nghe nhận 2 luồng audio cùng lúc:

Giọng gốc của người nói (qua WebRTC bình thường)
Giọng AI đọc bản dịch (luồng riêng)

Cần quyết định mix như thế nào — có 3 cách:

Ducking: giảm giọng gốc xuống 20-30%, tăng giọng dịch lên — nghe được cả hai
Replace: tắt hẳn giọng gốc, chỉ nghe giọng dịch — tự nhiên hơn nhưng mất ngữ điệu gốc
Người dùng tự chọn: slider điều chỉnh tỉ lệ mix theo sở thích
================================================
Voice-to-Subtitle đơn giản hơn Voice-to-Voice — bỏ hoàn toàn bước TTS, thay vào đó hiển thị chữ lên màn hình.

Luồng hoạt động
Bước 1 → 3 giống hệt Voice-to-Voice:

Thu âm → STT (nhận diện giọng) → Dịch văn bản
Bước 4 — Khác biệt hoàn toàn:

Thay vì chuyển thành giọng nói, văn bản đã dịch được push qua WebSocket xuống client và hiển thị dạng subtitle overlay phía dưới video — giống phụ đề phim.

Tại sao nhanh hơn Voice-to-Voice?
Bỏ được bước TTS (~200-400ms), nên tổng độ trễ chỉ còn ~300-600ms thay vì ~1s.
Ngoài ra text rất nhẹ (vài chục bytes), push qua WebSocket gần như tức thì — trong khi audio file nặng hơn nhiều và cần stream.

Người dùng thấy gì?
Mỗi người tự chọn ngôn ngữ subtitle trong settings. Ví dụ trong cùng một meeting:

Người Việt chọn xem subtitle tiếng Việt
Người Nhật chọn subtitle tiếng Nhật
Người Anh không bật subtitle vì nghe được trực tiếp

Server dịch song song ra nhiều ngôn ngữ cùng lúc từ một luồng STT, push riêng cho từng người.

Hạn chế so với Voice-to-Voice
Phải nhìn màn hình liên tục để đọc — không tiện khi đang ghi chép hay nhìn chỗ khác. Voice-to-Voice giải quyết vấn đề này vì chỉ cần nghe.