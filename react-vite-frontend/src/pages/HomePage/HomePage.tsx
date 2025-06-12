// src/pages/HomePage/HomePage.tsx
import React, { useState, useEffect, useCallback } from 'react';
import styles from './HomePage.module.css';
import InvitationCard from '../../components/InvitationCard';
import AuthModal from '../../components/AuthModal';
import type { Invitation, AuthMode, User, Comment as CommentType } from '../../types';

// --- MOCK DATA (Giữ nguyên hoặc thay thế bằng API calls) ---
const PLACEHOLDER_AVATAR_USER = 'https://gravatar.com/avatar/197593169c82a923a28974d9805686bf?s=400&d=robohash&r=x';
const PLACEHOLDER_AVATAR_COMMENT = 'https://gravatar.com/avatar/197593169c82a923a28974d9805686bf?s=400&d=robohash&r=x';

const MOCK_INVITATIONS_DATA: Invitation[] = [
  {
    id: '1', userId: 'userA', hostName: 'Anh Ba Hưng', avatarUrl: PLACEHOLDER_AVATAR_USER, location: 'Quán Nhậu Bờ Kè Xóm Chiếu', time: '19:00 - Thứ Sáu, 24/05', createdAt: new Date(Date.now() - 2 * 60 * 60 * 1000).toISOString(), venueImageUrl: 'https://pasgo.vn/Upload/anh-kham-pha/review-20--quan-nhau-view-dep-o-tphcm-duoc-yeu-thich---xem-ngay-uu-dai-155484781405.webp', description: 'Cuối tuần rồi, làm vài chai giải sầu anh em ơi! Kèo này bao vui, có mồi ngon nhạc xập xình.', participants: 3, likes: ['userB', 'userC'],
    comments: [ { id: 'c1', userId: 'userB', userName: 'Bé Tư', userAvatarUrl: PLACEHOLDER_AVATAR_COMMENT, content: 'Em một suất nha anh Ba!', timestamp: new Date(Date.now() - 1 * 60 * 60 * 1000).toISOString() }, { id: 'c2', userId: 'userD', userName: 'Chú Năm', userAvatarUrl: PLACEHOLDER_AVATAR_COMMENT, content: 'Kèo này ngon nè, để chú xem lịch.', timestamp: new Date(Date.now() - 0.5 * 60 * 60 * 1000).toISOString() } ]
  },
  {
    id: '2', userId: 'userE', hostName: 'Chị Tư Giao', avatarUrl: PLACEHOLDER_AVATAR_USER, location: 'Lẩu Dê Đồng Xanh, Q.7', time: '18:30 - Thứ Bảy, 25/05', createdAt: new Date(Date.now() - 25 * 60 * 60 * 1000).toISOString(), venueImageUrl: 'https://bloganchoi.com/wp-content/uploads/2022/04/quan-nhau-tai-tphcm-rong-rai-696x522.jpg', description: 'Kèo lẩu dê nóng hổi, trời mưa lành lạnh ăn là bá cháy. Ai tham gia điểm danh sớm nha.', likes: ['userA'], comments: []
  },
  {
    id: '3', userId: 'userF', hostName: 'Bạn Tèo', avatarUrl: PLACEHOLDER_AVATAR_USER, location: 'Ốc Đêm Bình Thạnh', time: '21:00 - Chủ Nhật, 26/05', createdAt: new Date(Date.now() - 5 * 24 * 60 * 60 * 1000).toISOString(), description: 'Thèm ốc quá ae ơi, làm vài dĩa lai rai cuối tuần nào. Quán này ốc tươi, giá phải chăng.', likes: ['userA', 'userB', 'userC', 'userD'],
    comments: [ { id: 'c3', userId: 'userG', userName: 'My', userAvatarUrl: PLACEHOLDER_AVATAR_COMMENT, content: 'Đi nè Tèo ơi!', timestamp: new Date(Date.now() - 4 * 24 * 60 * 60 * 1000).toISOString() } ]
  },
];
// --- END MOCK DATA ---


const HomePage: React.FC = () => {
  const [invitations, setInvitations] = useState<Invitation[]>([]);
  const [showAuthModal, setShowAuthModal] = useState(false);
  const [authModalMode, setAuthModalMode] = useState<AuthMode>('login');
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [selectedInvitationId, setSelectedInvitationId] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true); // Thêm state loading ban đầu

  // Hàm fetch thông tin user dựa trên token (giả lập)
  const fetchUserProfile = useCallback(async (token: string): Promise<User | null> => {
    // TODO: Gọi API backend của bạn để xác thực token và lấy thông tin user
    // Ví dụ: const response = await fetch('/api/v1/me', { headers: { 'Authorization': `Bearer ${token}` } });
    // if (!response.ok) return null;
    // const userData = await response.json();
    // return userData;

    console.log("Fetching user profile with token:", token); // Log để debug
    // Giả lập: Nếu có token, trả về một user giả định (trong thực tế, bạn sẽ gọi API)
    if (token === "mock_valid_token_userC") { // Thay "mock_valid_token_userC" bằng cách bạn lấy token thực
      return {
        id: 'userC', // Hoặc account_id
        username: 'UserTừToken', // Tên user thực từ API
        email: 'tokenuser@example.com',
        avatarUrl: PLACEHOLDER_AVATAR_USER,
        token: token, // Giữ lại token
      };
    }
    return null;
  }, []);


  // Load invitations và kiểm tra trạng thái đăng nhập khi component mount
  useEffect(() => {
    const loadInitialData = async () => {
      setIsLoading(true);
      // TODO: Fetch invitations từ API
      // const fetchedInvitations = await api.getInvitations();
      // setInvitations(fetchedInvitations);
      setInvitations(MOCK_INVITATIONS_DATA); // Tạm thời dùng mock

      const token = localStorage.getItem('authToken');
      if (token) {
        const userProfile = await fetchUserProfile(token);
        if (userProfile) {
          setCurrentUser(userProfile);
        } else {
          // Token không hợp lệ hoặc hết hạn, xóa token
          localStorage.removeItem('authToken');
        }
      }
      setIsLoading(false);
    };
    loadInitialData();
  }, [fetchUserProfile]);


  const handleJoinClick = (invitationId: string) => {
    if (!currentUser) {
      setSelectedInvitationId(invitationId);
      setAuthModalMode('login');
      setShowAuthModal(true);
    } else {
      alert(`Đã đăng nhập! Tham gia kèo ${invitationId}! (TODO: Implement join logic, gửi token: ${currentUser.token})`);
      // TODO: Gọi API để tham gia, gửi currentUser.token trong header Authorization
    }
  };

  const handleAuthSuccess = (loggedInUser: User) => {
    setCurrentUser(loggedInUser);
    if (loggedInUser.token) {
      localStorage.setItem('authToken', loggedInUser.token);
    }
    setShowAuthModal(false);
    alert(`Chào mừng ${loggedInUser.username || loggedInUser.username}!`);

    if (selectedInvitationId) {
      // Xử lý logic tham gia kèo ngay sau khi đăng nhập thành công
      handleJoinClick(selectedInvitationId); // Gọi lại handleJoinClick để nó chạy với currentUser mới
      setSelectedInvitationId(null);
    }
  };

  const handleLogout = () => {
    setCurrentUser(null);
    localStorage.removeItem('authToken');
    alert('Bạn đã đăng xuất.');
  };

  const handleLike = (invitationId: string) => {
    if (!currentUser) {
      setAuthModalMode('login');
      setShowAuthModal(true);
      setSelectedInvitationId(invitationId); // Lưu lại để có thể like sau khi login
      return;
    }
    setInvitations(prevInvitations =>
      prevInvitations.map(inv => {
        if (inv.id === invitationId) {
          const userHasLiked = inv.likes.includes(currentUser.id);
          const newLikes = userHasLiked
            ? inv.likes.filter(userId => userId !== currentUser.id)
            : [...inv.likes, currentUser.id];
          return { ...inv, likes: newLikes };
        }
        return inv;
      })
    );
    // TODO: Gọi API để cập nhật like, gửi currentUser.token
    console.log(`Liked invitation ${invitationId} with token ${currentUser.token}`);
  };

  const handleCommentSubmit = (invitationId: string, commentText: string) => {
    if (!currentUser) {
      setAuthModalMode('login');
      setShowAuthModal(true);
      setSelectedInvitationId(invitationId); // Lưu lại để có thể comment sau khi login
      return;
    }
    const newComment: CommentType = {
      id: `c${Date.now()}`,
      userId: currentUser.id,
      userName: currentUser.username || currentUser.username, // Sử dụng username nếu name không có
      userAvatarUrl: currentUser.avatarUrl || PLACEHOLDER_AVATAR_COMMENT,
      content: commentText,
      timestamp: new Date().toISOString(),
    };
    setInvitations(prevInvitations =>
      prevInvitations.map(inv => {
        if (inv.id === invitationId) {
          return { ...inv, comments: [...inv.comments, newComment] };
        }
        return inv;
      })
    );
    // TODO: Gọi API để gửi comment, gửi currentUser.token
    console.log(`Commented on invitation ${invitationId} with token ${currentUser.token}: ${commentText}`);
  };

  const openLoginModal = () => {
    setAuthModalMode('login');
    setShowAuthModal(true);
  };

  const openRegisterModal = () => {
    setAuthModalMode('register');
    setShowAuthModal(true);
  };

  if (isLoading) {
    return <div className={styles.loadingScreen}>Đang tải dữ liệu...</div>; // Hoặc một spinner đẹp hơn
  }

  return (
    <div className={styles.pageContainer}>
      <header className={styles.header}>
        <h1 className={styles.logo}>ZoNhau</h1>
        {currentUser ? (
          <div className={styles.userInfo}>
            <img src={currentUser.avatarUrl || PLACEHOLDER_AVATAR_USER} alt={currentUser.username || currentUser.username} className={styles.headerAvatar}/>
            <span>Chào, {currentUser.username || currentUser.username}!</span>
            <button onClick={handleLogout} className={`${styles.actionButton} ${styles.logoutButton}`}>Đăng Xuất</button>
          </div>
        ) : (
          <div className={styles.authActions}>
             <button onClick={openLoginModal} className={styles.actionButton}>
                Đăng nhập
             </button>
             <span className={styles.authSeparator}>/</span>
             <button onClick={openRegisterModal} className={`${styles.actionButton} ${styles.registerButton}`}>
                Đăng ký
             </button>
          </div>
        )}
      </header>

      {currentUser && (
        <div className={styles.createPostPlaceholder}>
          <img src={currentUser.avatarUrl || PLACEHOLDER_AVATAR_USER} alt="Bạn" className={styles.headerAvatar} />
          <input type="text" placeholder={`Hôm nay rủ rê gì đây, ${currentUser.username || currentUser.username}?`} readOnly onClick={() => alert("TODO: Mở form tạo lời mời")} />
        </div>
      )}

      <main className={styles.feedContainer}>
        {invitations.length === 0 ? (
          <p className={styles.emptyFeed}>Chưa có lời mời nào. Hãy là người đầu tiên tạo kèo!</p>
        ) : (
          invitations.map((inv) => (
            <InvitationCard
              key={inv.id}
              invitation={inv}
              currentUser={currentUser}
              onJoinClick={handleJoinClick}
              onLikeClick={handleLike}
              onCommentSubmit={handleCommentSubmit}
            />
          ))
        )}
      </main>

      <AuthModal
        isOpen={showAuthModal}
        onClose={() => { setShowAuthModal(false); setSelectedInvitationId(null); setAuthModalMode('login'); }}
        initialMode={authModalMode}
        onAuthSuccess={handleAuthSuccess}
      />
    </div>
  );
}

export default HomePage;