// src/components/AuthModal/AuthModal.tsx
import React, { useState, useEffect } from 'react';
import LoginForm from '../LoginForm';
import RegisterForm from '../RegisterForm';
import type { AuthMode, User } from '../../types';
import styles from './AuthModal.module.css';

interface AuthModalProps {
  isOpen: boolean;
  onClose: () => void;
  initialMode?: AuthMode;
  onAuthSuccess: (userData: User) => void;
}

const AuthModal: React.FC<AuthModalProps> = ({
  isOpen,
  onClose,
  initialMode = 'login',
  onAuthSuccess,
}) => {
  const [mode, setMode] = useState<AuthMode>(initialMode);
  const [authError, setAuthError] = useState<string | null>(null);

  useEffect(() => {
    if (isOpen) {
      setMode(initialMode);
      setAuthError(null);
    }
  }, [initialMode, isOpen]);

  if (!isOpen) return null;

  // SỬA Ở ĐÂY: đưa "=>" lên cùng dòng hoặc ngay sau ")"
  const handleLoginSuccess = (userData: User): void => { // Hoặc bỏ :void cũng được, TS sẽ tự suy luận
    console.log('AuthModal: Login successful', userData);
    setAuthError(null);
    onAuthSuccess(userData);
  };

  // SỬA Ở ĐÂY: tương tự
  const handleRegisterSuccess = (registrationData: any): void => {
    console.log('AuthModal: Registration successful', registrationData);
    setAuthError(null);
    alert('Đăng ký thành công! Vui lòng đăng nhập để tiếp tục.');
    setMode('login');
    // ... (phần logic tự động login nếu có)
  };

  const handleFormError = (errorMessage: string): void => {
    setAuthError(errorMessage);
  };

  return (
    <div className={styles.modalOverlay} onClick={onClose}>
      <div className={styles.modalContent} onClick={(e) => e.stopPropagation()}>
        <button className={styles.modalCloseButton} onClick={onClose}>×</button>
        {authError && <p className={styles.modalErrorMessage}>{authError}</p>}

        {mode === 'login' ? (
          <LoginForm
            onLoginSuccess={handleLoginSuccess}
            onSwitchToRegister={() => { setMode('register'); setAuthError(null); }}
            onLoginError={handleFormError}
          />
        ) : (
          <RegisterForm
            onRegisterSuccess={handleRegisterSuccess}
            onSwitchToLogin={() => { setMode('login'); setAuthError(null); }}
            onRegisterError={handleFormError}
          />
        )}
      </div>
    </div>
  );
};

export default AuthModal;