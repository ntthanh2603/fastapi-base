
import React, { useState } from 'react';
import styles from './RegisterForm.module.css';

interface RegisterFormProps {
  // onRegisterSuccess sẽ được gọi khi backend trả về thành công
  onRegisterSuccess: (registrationData: any) => void; // Thay 'any' bằng type cụ thể nếu backend trả về user data
  onSwitchToLogin: () => void;
  // Tùy chọn: hàm xử lý lỗi từ backend
  onRegisterError?: (errorMessage: string) => void;
}

const RegisterForm: React.FC<RegisterFormProps> = ({
  onRegisterSuccess,
  onSwitchToLogin,
  onRegisterError,
}) => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null); // Xóa lỗi cũ

    if (!username.trim()) {
      setError('Vui lòng nhập tên tài khoản.');
      return;
    }
    if (!email.trim()) {
      setError('Vui lòng nhập email.');
      return;
    }
    // Thêm validate email cơ bản
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
        setError('Địa chỉ email không hợp lệ.');
        return;
    }
    if (!password) {
      setError('Vui lòng nhập mật khẩu.');
      return;
    }
    if (password.length < 6) { // Giả sử mật khẩu tối thiểu 6 ký tự
        setError('Mật khẩu phải có ít nhất 6 ký tự.');
        return;
    }
    if (password !== confirmPassword) {
      setError('Mật khẩu xác nhận không khớp!');
      return;
    }

    setIsLoading(true);

    try {
      const response = await fetch('http://localhost:8800/v1/2025/auth/register', { // Sử dụng HTTP
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          username: username.trim(),
          email: email.trim(),
          password: password,
          lang: 'vi' 
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        // Xử lý lỗi từ backend (ví dụ: username/email đã tồn tại)
        const errorMessage = data.message || `Đăng ký thất bại (HTTP ${response.status})`;
        setError(errorMessage);
        if (onRegisterError) {
          onRegisterError(errorMessage);
        }
        return;
      }

      // Đăng ký thành công
      console.log('Register successful:', data);
      onRegisterSuccess(data); // Truyền dữ liệu trả về từ backend (có thể là thông tin user hoặc token)

    } catch (err) {
      console.error('Registration error:', err);
      const errorMessage = 'Đã có lỗi xảy ra trong quá trình đăng ký. Vui lòng thử lại.';
      setError(errorMessage);
      if (onRegisterError) {
        onRegisterError(errorMessage);
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className={styles.authForm}>
      <h2>Đăng Ký Tài Khoản</h2>
      {error && <p className={styles.errorMessage}>{error}</p>}
      <div className={styles.formGroup}>
        <label htmlFor="register-username">Tên tài khoản:</label>
        <input
          type="text"
          id="register-username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Ví dụ: anhtuan123"
          disabled={isLoading}
          // required // HTML5 validation, nhưng chúng ta đã validate bằng JS
        />
      </div>
      <div className={styles.formGroup}>
        <label htmlFor="register-email">Email:</label>
        <input
          type="email"
          id="register-email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="vidu@email.com"
          disabled={isLoading}
          // required
        />
      </div>
      <div className={styles.formGroup}>
        <label htmlFor="register-password">Mật khẩu:</label>
        <input
          type="password"
          id="register-password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Ít nhất 6 ký tự"
          disabled={isLoading}
          // required
        />
      </div>
      <div className={styles.formGroup}>
        <label htmlFor="confirm-password">Xác nhận mật khẩu:</label>
        <input
          type="password"
          id="confirm-password"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
          placeholder="Nhập lại mật khẩu"
          disabled={isLoading}
          // required
        />
      </div>
      <button type="submit" className={styles.submitButton} disabled={isLoading}>
        {isLoading ? 'Đang xử lý...' : 'Đăng Ký'}
      </button>
      <p className={styles.switchToLogin}>
        Đã có tài khoản?{' '}
        <button type="button" onClick={onSwitchToLogin} className={styles.linkButton} disabled={isLoading}>
          Đăng nhập ngay
        </button>
      </p>
    </form>
  );
};

export default RegisterForm;