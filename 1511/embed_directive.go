package main

import "embed" //cho phép bạn nhúng các tệp và thư mục từ hệ thống tệp vào bên trong ứng dụng
//hữu ích khi muốn đóng gói các tệp văn bản, HTML, hoặc dữ liệu tĩnh khác vào ứng dụng của mình mà không cần cung cấp các tệp bên ngoài

//go:embed folder/single_file.txt
var fileString string

//Dòng này nhúng tệp folder/single_file.txt vào biến ngay bên dưới nó

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

//Các dòng này nhúng tất cả các tệp có phần mở rộng .hash trong thư mục folder (bao gồm cả single_file.txt) vào biến folder, kiểu embed.FS

func main() {
	print(fileString)
	print(string(fileByte))
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

//Trước hết, phải import embed vào bất kỳ tệp nào sử dụng chỉ thị nhúng.
//Nếu cần nhập nhúng nhưng không tham chiếu đến bất kỳ mã định danh nào được xuất trong đó,
//nên sử dụng import _ "embed"để yêu cầu Go nhập nhúng ngay cả khi có vẻ như nó không được sử dụng
//Thứ hai, chỉ có thể sử dụng //go:embed cho các biến ở cấp độ gói, không phải trong các hàm hoặc phương thức
//Thứ ba, khi bao gồm một thư mục, nó sẽ không bao gồm các tệp bắt đầu bằng .hoặc _, nhưng nếu sử dụng ký tự đại diện,
//như dir/*, nó sẽ bao gồm tất cả các tệp khớp, ngay cả khi chúng bắt đầu bằng .hoặc _.
//Hãy lưu ý rằng việc vô tình bao gồm .DS_Storecác tệp của Mac OS có thể là vấn đề bảo mật trong trường hợp muốn nhúng
//các tệp vào máy chủ web nhưng không cho phép người dùng xem danh sách tất cả các tệp.
//Vì lý do này, việc sử dụng //go:embed dir/*gần như luôn là một lỗi . Hãy sử dụng //go:embed dirhoặc //go:embed dir/*.extkhi cần thay thế
