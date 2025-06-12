// src/types/index.ts
import type { Locale } from 'date-fns'; // Thêm nếu dùng date-fns

export interface Comment {
  id: string;
  userId: string;
  userName: string;
  userAvatarUrl?: string; // Thêm avatar cho người bình luận
  content: string;
  timestamp: string; // ISO String
}

export interface Invitation {
  id: string;
  userId: string;
  hostName: string;
  avatarUrl?: string; // Avatar của người host
  location: string;
  time: string; // Thời gian diễn ra sự kiện
  createdAt: string; // Thời gian đăng bài (ISO String) - QUAN TRỌNG
  venueImageUrl?: string; // Hình ảnh của quán/địa điểm - QUAN TRỌNG
  description: string;
  participants?: number;
  likes: string[];
  comments: Comment[];
}

export interface User {
  id: string; // Hoặc account_id từ bảng drunk_user
  username: string;
  email: string;
  avatarUrl?: string;
  is_vip?: boolean; // Tùy chọn, dựa trên bảng
  // Thêm các trường khác nếu backend trả về sau khi login/register
  token?: string; // QUAN TRỌNG: Nếu backend trả về token JWT
}

export type AuthMode = 'login' | 'register';

// Thêm interface cho date-fns locale nếu bạn muốn hỗ trợ tiếng Việt
export interface DateFnsLocale {
  locale: Locale;
  // Bất kỳ thuộc tính nào khác mà bạn cần cho locale
}