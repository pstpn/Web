import AuthService from './auth.service';
import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';

const mock = new MockAdapter(axios);
const API_URL = 'http://localhost:8081/api/v1/';

describe('auth.service', () => {
  afterEach(() => {
    mock.reset();
    localStorage.clear();
  });

  it('should login user and save user data to localStorage', async () => {
    const user = {phoneNumber: '1234567890', password: 'password'};
    const mockResponse = {accessToken: 'fake-jwt-access-token', refreshToken: 'fake-jwt-refresh-token'};

    mock.onPost(`${API_URL}login`).reply(200, mockResponse);

    const data = await AuthService.login(user);

    expect(data).toEqual(mockResponse);
  });

  it('should logout user and remove user data from localStorage', () => {
    localStorage.setItem('user', JSON.stringify({accessToken: 'fake-jwt-access-token', refreshToken: 'fake-jwt-refresh-token'}));

    AuthService.logout();

    expect(localStorage.getItem('user')).toBeNull();
  });

  it('should register user and save user data to localStorage', async () => {
    const user = {
      phoneNumber: '1234567890',
      name: 'John',
      surname: 'Doe',
      selectedCompany: 0,
      post: 'Developer',
      password: 'password',
      dateOfBirth: '1990-01-01',
    };
    const mockResponse = {accessToken: 'fake-jwt-access-token', refreshToken: 'fake-jwt-refresh-token'};

    mock.onPost(`${API_URL}register`).reply(200, mockResponse);

    const data = await AuthService.register(user);

    expect(data).toEqual(mockResponse);
  });

  it('should refresh tokens and update localStorage', async () => {
    const user = {
      accessToken: 'oldAccessToken',
      refreshToken: 'oldRefreshToken',
    };
    const mockResponse = {accessToken: 'newAccessToken', refreshToken: 'newRefreshToken'};

    mock.onPost(`${API_URL}refresh`).reply(200, mockResponse);

    const data = await AuthService.refreshTokens(user);

    expect(data).toEqual(mockResponse);
  });
});
