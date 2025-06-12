// src/components/InvitationCard/InvitationCard.tsx
import React, { useState } from 'react';
import type { Invitation, User, Comment as CommentType } from '../../types';
import styles from './InvitationCard.module.css';
import { formatDistanceToNow, parseISO } from 'date-fns';
import { vi } from 'date-fns/locale'; // Import tiếng Việt

// Sử dụng lại các placeholder nếu cần
const PLACEHOLDER_AVATAR_USER_CARD = 'https://via.placeholder.com/40?text=H';
const PLACEHOLDER_AVATAR_COMMENT_CARD = 'https://via.placeholder.com/32?text=C';


interface InvitationCardProps {
  invitation: Invitation;
  currentUser: User | null;
  onJoinClick: (invitationId: string) => void;
  onLikeClick: (invitationId: string) => void;
  onCommentSubmit: (invitationId: string, commentText: string) => void;
}

const InvitationCard: React.FC<InvitationCardProps> = ({
  invitation,
  currentUser,
  onJoinClick,
  onLikeClick,
  onCommentSubmit,
}) => {
  const [showComments, setShowComments] = useState(false);
  const [newComment, setNewComment] = useState('');

  const userHasLiked = currentUser && invitation.likes.includes(currentUser.id);

  const handleCommentSubmitForm = (e: React.FormEvent) => {
    e.preventDefault();
    if (newComment.trim() && currentUser) {
      onCommentSubmit(invitation.id, newComment.trim());
      setNewComment('');
      // setShowComments(true); // Giữ nguyên trạng thái comments
    } else if (!currentUser) {
        alert('Bạn cần đăng nhập để bình luận!');
    }
  };

  const formatRelativeTime = (isoDateString: string) => {
    try {
      const date = parseISO(isoDateString);
      return formatDistanceToNow(date, { addSuffix: true, locale: vi });
    } catch (error) {
      console.error("Invalid date for relative time:", isoDateString);
      return invitation.time; // Fallback to event time or a default string
    }
  };

  return (
    <div className={styles.card}>
      <div className={styles.cardHeader}>
        <img
          src={invitation.avatarUrl || PLACEHOLDER_AVATAR_USER_CARD}
          alt={invitation.hostName}
          className={styles.avatar}
        />
        <div className={styles.hostInfo}>
          <span className={styles.hostName}>{invitation.hostName}</span>
          {/* Thời gian đăng bài */}
          <span className={styles.postTime}>{formatRelativeTime(invitation.createdAt)}</span>
        </div>
      </div>

      <div className={styles.cardContent}>
        <p className={styles.description}>{invitation.description}</p>
        {/* Hình ảnh của quán */}
        {invitation.venueImageUrl && (
          <img src={invitation.venueImageUrl} alt={`Địa điểm: ${invitation.location}`} className={styles.venueImage} />
        )}
        <p className={styles.eventDetails}>
            <strong>Địa điểm:</strong> {invitation.location}
        </p>
        <p className={styles.eventDetails}>
            <strong>Thời gian diễn ra:</strong> {invitation.time}
        </p>
        {invitation.participants !== undefined && (
          <p className={styles.eventDetails}><em>Số người đã tham gia: {invitation.participants}</em></p>
        )}
      </div>

      <div className={styles.cardStats}>
        <span>{invitation.likes.length} Lượt thích</span>
        <span onClick={() => setShowComments(!showComments)} className={styles.commentToggle}>
          {invitation.comments.length} Bình luận
        </span>
      </div>

      <div className={styles.cardActions}>
        <button
          className={`${styles.actionButton} ${userHasLiked ? styles.liked : ''}`}
          onClick={() => onLikeClick(invitation.id)}
          disabled={!currentUser}
        >
          {/* Icon Thích (SVG hoặc FontAwesome) có thể thêm vào đây */}
          {userHasLiked ? 'Đã thích' : 'Thích'}
        </button>
        <button className={styles.actionButton} onClick={() => setShowComments(prev => !prev)} /*disabled={!currentUser}*/>
          {/* Icon Bình luận */}
          Bình luận
        </button>
        <button className={styles.actionButton} onClick={() => onJoinClick(invitation.id)}>
          {/* Icon Tham gia */}
          Tham Gia
        </button>
      </div>

      {showComments && (
        <div className={styles.commentsSection}>
          {invitation.comments.length > 0 ? (
            invitation.comments.slice(0,3).map((comment) => ( // Hiển thị 3 comment mới nhất ví dụ
              <div key={comment.id} className={styles.comment}>
                <img src={comment.userAvatarUrl || PLACEHOLDER_AVATAR_COMMENT_CARD} alt={comment.userName} className={styles.commentAvatar} />
                <div className={styles.commentBody}>
                  <span className={styles.commentUser}>{comment.userName}</span>
                  <p className={styles.commentText}>{comment.content}</p>
                  <span className={styles.commentTime}>{formatRelativeTime(comment.timestamp)}</span>
                </div>
              </div>
            ))
          ) : (
            <p className={styles.noComments}>Chưa có bình luận nào.</p>
          )}
          {invitation.comments.length > 3 && (
            <p className={styles.viewMoreComments} onClick={() => alert("TODO: Xem tất cả bình luận")}>Xem thêm bình luận...</p>
          )}

          {currentUser && (
            <form onSubmit={handleCommentSubmitForm} className={styles.commentForm}>
              <img src={currentUser.avatarUrl || PLACEHOLDER_AVATAR_COMMENT_CARD} alt="Bạn" className={styles.commentAvatar} />
              <input
                type="text"
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                placeholder="Viết bình luận..."
                className={styles.commentInput}
              />
              <button type="submit" className={styles.commentSubmitButton} disabled={!newComment.trim()}>Gửi</button>
            </form>
          )}
        </div>
      )}
    </div>
  );
};

export default InvitationCard;