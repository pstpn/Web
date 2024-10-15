import {auth} from './auth.module';
import AuthService from '../api/auth.service';

jest.mock('../api/auth.service');

describe('auth.module', () => {
  let commit;

  beforeEach(() => {
    commit = jest.fn();
  });

  describe('actions', () => {
    it('login action should commit loginSuccess on success', async () => {
      const user = {phoneNumber: '1234567890', password: 'password'};
      const mockResponse = {id: 1, fullName: 'John Doe', accessToken: 'fake-jwt-access-token'};
      AuthService.login.mockResolvedValue(mockResponse);

      await auth.actions.login({commit}, user);

      expect(commit).toHaveBeenCalledWith('loginSuccess', mockResponse);
    });

    it('login action should commit loginFailure on error', async () => {
      const user = {phoneNumber: '1234567890', password: 'password'};
      const mockError = new Error('Login failed');
      AuthService.login.mockRejectedValue(mockError);

      try {
        await auth.actions.login({commit}, user);
      } catch (error) {
        expect(commit).toHaveBeenCalledWith('loginFailure');
      }
    });

    it('logout action should commit logout', () => {
      auth.actions.logout({commit});

      expect(commit).toHaveBeenCalledWith('logout');
      expect(AuthService.logout).toHaveBeenCalled();
    });

    it('register action should commit registerSuccess on success', async () => {
      const user = {phoneNumber: '1234567890', password: 'password'};
      const mockResponse = {data: {message: 'Registration successful'}};
      AuthService.register.mockResolvedValue(mockResponse);

      await auth.actions.register({commit}, user);

      expect(commit).toHaveBeenCalledWith('registerSuccess');
    });

    it('register action should commit registerFailure on error', async () => {
      const user = {phoneNumber: '1234567890', password: 'password'};
      const mockError = new Error('Registration failed');
      AuthService.register.mockRejectedValue(mockError);

      try {
        await auth.actions.register({commit}, user);
      } catch (error) {
        expect(commit).toHaveBeenCalledWith('registerFailure');
      }
    });

    it('refreshTokens action should commit loginSuccess on success', async () => {
      const user = {accessToken: 'oldAccessToken', refreshToken: 'oldRefreshToken'};
      const mockResponse = {data: {accessToken: 'newAccessToken', refreshToken: 'newRefreshToken'}};
      AuthService.refreshTokens.mockResolvedValue(mockResponse);

      await auth.actions.refreshTokens({commit}, user);

      expect(commit).toHaveBeenCalledWith('loginSuccess', user);
    });

    it('refreshTokens action should commit loginFailure on error', async () => {
      const user = {accessToken: 'oldAccessToken', refreshToken: 'oldRefreshToken'};
      const mockError = new Error('Token refresh failed');
      AuthService.refreshTokens.mockRejectedValue(mockError);

      try {
        await auth.actions.refreshTokens({commit}, user);
      } catch (error) {
        expect(commit).toHaveBeenCalledWith('loginFailure');
      }
    });
  });

  describe('mutations', () => {
    it('loginSuccess mutation should update state correctly', () => {
      const state = {status: {loggedIn: false}, user: null};
      const user = {id: 1, fullName: 'John Doe', accessToken: 'fake-jwt-access-token'};

      auth.mutations.loginSuccess(state, user);

      expect(state.status.loggedIn).toBe(true);
      expect(state.user).toEqual(user);
    });

    it('loginFailure mutation should update state correctly', () => {
      const state = {status: {loggedIn: true}, user: {id: 1, fullName: 'John Doe'}};

      auth.mutations.loginFailure(state);

      expect(state.status.loggedIn).toBe(false);
      expect(state.user).toBeNull();
    });

    it('logout mutation should reset state', () => {
      const state = {status: {loggedIn: true}, user: {id: 1, fullName: 'John Doe'}};

      auth.mutations.logout(state);

      expect(state.status.loggedIn).toBe(false);
      expect(state.user).toBeNull();
    });

    it('registerSuccess mutation should update state correctly', () => {
      const state = {status: {loggedIn: false}};

      auth.mutations.registerSuccess(state);

      expect(state.status.loggedIn).toBe(false);
    });

    it('registerFailure mutation should update state correctly', () => {
      const state = {status: {loggedIn: false}};

      auth.mutations.registerFailure(state);

      expect(state.status.loggedIn).toBe(false);
    });
  });
});
