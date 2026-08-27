const username = document.getElementById("username");
const password = document.getElementById("password");
const registerbtn = document.getElementById("registerbtn");
const loginbtn = document.getElementById("loginbtn");

registerbtn.addEventListener('click', function () {
  const LoginData = {
    username: username.value,
    password: password.value
  };

  RegisterUser(LoginData);
});

loginbtn.addEventListener('click', function () {
  const LoginData = {
    username: username.value,
    password: password.value
  };

  LoginUser(LoginData);
});

async function RegisterUser(LoginData) {
  try {
    const response = await fetch('api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(LoginData)
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error('Registration failed');
    }

    console.log('Registration successful:', data);
    window.location.href = '/';

  } catch (error) {
    console.error('Registration error:', error);
  }
}

async function LoginUser(LoginData) {
  try {
    const response = await fetch('api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include',
      body: JSON.stringify(LoginData)
    });

    const data = await response.json();
    if (!response.ok) {
      throw new Error('Login failed');
    }

    console.log('Login successful');
    window.location.href = data.redirect || '/home';
  } catch (error) {
    console.error('Login error:', error);
  }
}
