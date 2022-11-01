/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `menu_items` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `text` longtext,
  `link` longtext,
  `id_parent` bigint DEFAULT NULL,
  `order_number` bigint DEFAULT NULL,
  `level` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `menu_items` (`id`, `text`, `link`, `id_parent`, `order_number`, `level`) VALUES
(1, 'Shop', NULL, 0, 1, 1);
INSERT INTO `menu_items` (`id`, `text`, `link`, `id_parent`, `order_number`, `level`) VALUES
(2, 'Tranh Canvas', NULL, 0, 2, 1);
INSERT INTO `menu_items` (`id`, `text`, `link`, `id_parent`, `order_number`, `level`) VALUES
(3, 'Tất Cả Chủ Đề', NULL, 2, 1, 2);
INSERT INTO `menu_items` (`id`, `text`, `link`, `id_parent`, `order_number`, `level`) VALUES
(4, 'Loài Hoa', NULL, 2, 2, 2),
(5, 'Tất Cả', NULL, 4, 1, 3),
(6, 'Hoa Hồng', NULL, 4, 2, 3),
(7, 'Hoa Mẫu Đơn', NULL, 4, 3, 3),
(8, 'Hoa Sen', NULL, 4, 4, 3),
(9, 'Hoa Lưu Ly', NULL, 4, 5, 3),
(10, 'Sen Đá', NULL, 4, 6, 3),
(11, 'Hoa Tulip', NULL, 4, 7, 3),
(12, 'Hướng Dương', NULL, 4, 8, 3),
(13, 'Hoa Lan', NULL, 4, 9, 3),
(14, 'Hoa Cúc', NULL, 4, 10, 3),
(15, 'Cây Hoa', NULL, 4, 11, 3),
(16, 'Bồ Công Anh', NULL, 4, 12, 3),
(17, 'Bình Hoa', NULL, 4, 13, 3),
(18, 'Cẩm Chướng', NULL, 4, 14, 3),
(19, 'Xương Rồng', NULL, 4, 15, 3),
(20, 'Hoa Khác', NULL, 4, 16, 3),
(21, 'Lá Cây', NULL, 2, 3, 2),
(22, 'Tất Cả', NULL, 21, 1, 3),
(23, 'Kiểu Vẽ', NULL, 21, 2, 3),
(24, 'Kiểu Chụp', NULL, 21, 3, 3),
(25, 'Trừu Tượng, Tối Giản', NULL, 2, 4, 2),
(26, 'Tranh Trừu Tượng', NULL, 25, 1, 3),
(27, 'Tôn Giáo', NULL, 2, 5, 2),
(28, 'Tất Cả', NULL, 27, 1, 3),
(29, 'Thiên Chúa', NULL, 27, 2, 3),
(30, 'Phật Giáo', NULL, 27, 3, 3),
(31, 'Phong Cảnh', NULL, 2, 6, 2),
(32, 'Tất Cả', NULL, 31, 1, 3),
(33, 'Tranh Núi', NULL, 31, 2, 3),
(34, 'Tranh Rừng', NULL, 31, 3, 3),
(35, 'Tranh Biển', NULL, 31, 4, 3),
(36, 'Vườn Hoa', NULL, 31, 5, 3),
(37, 'Đồng Lúa', NULL, 31, 6, 3),
(38, 'Đồng Quê', NULL, 31, 7, 3),
(39, 'Thác Nước', NULL, 31, 8, 3),
(40, 'Phong Thủy', NULL, 2, 7, 2),
(41, 'Tất Cả', NULL, 40, 1, 3),
(42, 'Mã Đáo T.Công', NULL, 40, 2, 3),
(43, 'Thuận Buồm X.Gió', NULL, 40, 3, 3),
(44, 'Cửu Ngư', NULL, 40, 4, 3),
(45, 'Công Phượng', NULL, 40, 5, 3),
(46, 'Tứ Quý', NULL, 40, 6, 3),
(47, 'Tùng Hạc', NULL, 40, 7, 3),
(48, 'Cô Gái', NULL, 2, 8, 2),
(49, 'Tất Cả', NULL, 48, 1, 3),
(50, 'Cá Tính', NULL, 48, 2, 3),
(51, 'Dịu Dàng', NULL, 48, 3, 3),
(52, 'Quyến Rũ', NULL, 48, 4, 3),
(53, 'Chủ Đề Khác', NULL, 2, 9, 2),
(54, 'Động Vật', NULL, 53, 1, 3),
(55, 'Hình Học T.Giản', NULL, 53, 2, 3),
(56, 'Địa Danh', NULL, 53, 3, 3),
(57, 'Hoạt Hình', NULL, 53, 4, 3),
(58, 'Hoa Quả', NULL, 53, 5, 3),
(59, 'Cảnh Vật', NULL, 53, 6, 3),
(60, 'T.Phẩm Danh Họa', NULL, 53, 7, 3),
(61, 'Tranh Sơn Dầu', NULL, 0, 3, 1),
(62, 'Thảm Sofa', NULL, 0, 4, 1),
(63, 'Thảm Sofa', NULL, 62, 1, 2),
(64, 'Thảm Lông Dài', NULL, 62, 2, 2),
(65, 'Thảm Lông Ngắn', NULL, 62, 3, 2),
(66, 'Sản Phẩm Khác', NULL, 0, 5, 1),
(67, 'Tranh Sơn Dầu', NULL, 66, 1, 2),
(68, 'Tranh Dán Tường', NULL, 66, 2, 2),
(69, 'Khung Hình Dán Tường Siêu Nhẹ', NULL, 66, 3, 2),
(70, 'Giấy Dán Tường', NULL, 66, 4, 2),
(71, 'Thảm Sofa', NULL, 66, 5, 2),
(72, 'Khuyến Mãi', NULL, 0, 6, 1),
(73, 'Hot Deal 15%', NULL, 72, 1, 2),
(74, 'VC Deal 10%', NULL, 72, 2, 2),
(75, 'Gallery', NULL, 0, 7, 1),
(76, 'Reviews', NULL, 0, 8, 1),
(77, 'Liên Hệ', NULL, 0, 9, 1),
(78, 'Tranh Tối Giản', NULL, 25, 2, 3);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;