
import React, { useState } from 'react';
import styles from './LoginForm.module.css'; 
import type { User } from '../../types'; 

interface LoginFormProps {
  onLoginSuccess: (userData: User) => void; 
  onSwitchToRegister: () => void;
  onLoginError: (errorMessage: string) => void;
}

const LoginForm: React.FC<LoginFormProps> = ({
  onLoginSuccess,
  onSwitchToRegister,
  onLoginError,
}) => {
  const [identifier, setIdentifier] = useState(''); 
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null); 

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null); 
    onLoginError(""); 

    if (!identifier.trim()) {
      setError('Vui lòng nhập tên tài khoản hoặc email.');
      return;
    }
    if (!password) {
      setError('Vui lòng nhập mật khẩu.');
      return;
    }

    setIsLoading(true);

    try {
      // Thay đổi URL và body cho phù hợp với API đăng nhập của bạn
      const response = await fetch('http://localhost:8800/v1/2025/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          identifier: identifier.trim(), // Hoặc 'usernameOrEmail', 'loginId' tùy backend
          password: password,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        const errorMessage = data.message || `Đăng nhập thất bại (HTTP ${response.status})`;
        setError(errorMessage); // Hiển thị lỗi trên LoginForm
        onLoginError(errorMessage); // Gửi lỗi lên AuthModal (nếu muốn)
        return;
      }

      // Đăng nhập thành công, backend nên trả về User object và token
      console.log('Login successful:', data);
      // Giả sử data có dạng { user: User, token: string }
      if (data.user && data.token) {
        const loggedInUser: User = {
            ...data.user, // Các trường từ bảng drunk_user mà backend trả về
            token: data.token,
        };
        onLoginSuccess(loggedInUser);
      } else {
        // Nếu backend trả về cấu trúc khác
        const errorMessage = "Dữ liệu đăng nhập không hợp lệ từ server.";
        setError(errorMessage);
        onLoginError(errorMessage);
      }


    } catch (err) {
      console.error('Login error:', err);
      const errorMessage = 'Đã có lỗi xảy ra trong quá trình đăng nhập.';
      setError(errorMessage);
      onLoginError(errorMessage);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className={styles.authForm}> {/* Sử dụng class của LoginForm */}
      <h2>Đăng Nhập</h2>
      {error && <p className={styles.errorMessage}>{error}</p>} {/* Hiển thị lỗi của LoginForm */}
      <div className={styles.formGroup}>
        <label htmlFor="login-identifier">Tên tài khoản hoặc Email:</label>
        <input
          type="text"
          id="login-identifier"
          value={identifier}
          onChange={(e) => setIdentifier(e.target.value)}
          disabled={isLoading}
          required
        />
      </div>
      <div className={styles.formGroup}>
        <label htmlFor="login-password">Mật khẩu:</label>
        <input
          type="password"
          id="login-password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          disabled={isLoading}
          required
        />
      </div>
      <button type="submit" className={styles.submitButton} disabled={isLoading}>
        {isLoading ? 'Đang đăng nhập...' : 'Đăng Nhập'}
      </button>
      <p className={styles.switchToRegister}> {/* Đổi tên class cho rõ nghĩa */}
        Chưa có tài khoản?{' '}
        <button type="button" onClick={onSwitchToRegister} className={styles.linkButton} disabled={isLoading}>
          Đăng ký ngay
        </button>
      </p>
    </form>
  );
};

export default LoginForm;