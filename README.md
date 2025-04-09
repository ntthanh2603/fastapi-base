# 🚀 FastAPI Base Project 🚀

Welcome to the **FastAPI Base Project**! 🎉 This is a lightweight, high-performance starter template built with [FastAPI](https://fastapi.tiangolo.com/), a modern Python web framework. Whether you're building APIs, microservices, or just experimenting, this project has you covered! 🌟

---

## ✨ Features

- ⚡ **Blazing Fast**: Powered by FastAPI for asynchronous, high-speed performance.
- 🛠️ **Ready to Use**: Pre-configured with essential tools and structure.
- 📝 **API Documentation**: Auto-generated docs with Swagger UI & ReDoc.
- 🐍 **Pythonic**: Built with Python 3.9+ for clean, modern code.
- 🌍 **Scalable**: Perfect for small projects or large-scale applications.

---

## 🛠️ Tech Stack

- **FastAPI**: The core framework for building APIs. 🚀
- **Uvicorn**: Lightning-fast ASGI server. ⚡
- **Pydantic**: Data validation and settings management. ✅
- **Python**: Because who doesn’t love Python? 🐍

---

## 📦 Installation

Get started in just a few steps! ⏩

1. **Clone the repo**:

   ```bash
   git clone https://github.com/ntthanh2603/fastapi-base.git
   cd fastapi-base
   ```

2. **Set up a virtual environment**:

   ```bash
   python -m venv venv
   source venv/bin/activate  # On Windows: venv\Scripts\activate
   ```

3. **Install dependencies**:

   ```bash
   pip install -r requirements.txt
   ```

4. **Run the app**:

   ```bash
   uvicorn main:app --reload
   ```

5. Open your browser at `http://127.0.0.1:3000` and enjoy! 🌐

---

## 🌐 Endpoints

- **GET /**: Welcome message. 👋
- **GET /docs**: Interactive API docs (Swagger UI). 📚
- **GET /redoc**: Alternative API docs (ReDoc). 📖

---

## 📂 Project Structure

```
fastapi-base-project/
├── main.py           # 🚀 Entry point of the app
├── requirements.txt  # 📋 Dependencies
├── README.md         # 📝 You’re reading it!
└── venv/             # 🗂️ Virtual environment
```

---

## 🧑‍💻 Usage

Start building your API by adding routes to `main.py`. Here’s a quick example:

```python
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"message": "Hello, World! 🌍"}
```

Run the server and test it out! 🎯

---

## 🌟 Contributing

Contributions are welcome! 🙌 Fork the repo, make your changes, and submit a pull request. Let’s make this project even better together! 🤝

1. Fork it 🍴
2. Create your feature branch (`git checkout -b feature/awesome`) 🌿
3. Commit your changes (`git commit -m 'Add something awesome'`) 💾
4. Push to the branch (`git push origin feature/awesome`) ⬆️
5. Open a Pull Request 🚀

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details. 📝

---

## 💬 Contact

Have questions? Reach out via [email](mailto:tuanthanh2kk4@gmail.com) or open an issue! 📧

Happy coding! 🎉🚀
