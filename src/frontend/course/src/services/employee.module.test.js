import {employee} from './employee.module';
import EmployeeService from '../api/employee.service';

jest.mock('../api/employee.service');

describe('employee.module', () => {
  let commit;

  beforeEach(() => {
    commit = jest.fn();
  });

  describe('actions', () => {
    it('fillProfile action should commit setFilled on success', async () => {
      const formData = {fullName: 'John Doe', position: 'Developer'};
      const mockResponse = {data: {message: 'Profile filled successfully'}};
      EmployeeService.fillProfile.mockResolvedValue(mockResponse);

      const response = await employee.actions.fillProfile({commit}, formData);

      expect(commit).toHaveBeenCalledWith('setFilled', true);
      expect(response).toEqual(mockResponse.data);
    });

    it('fillProfile action should handle error', async () => {
      const formData = {fullName: 'John Doe', position: 'Developer'};
      const mockError = new Error('Fill profile failed');
      EmployeeService.fillProfile.mockRejectedValue(mockError);

      await expect(employee.actions.fillProfile({commit}, formData)).rejects.toThrow(mockError);
    });

    it('getProfile action should commit setProfile on success', async () => {
      const mockProfile = {id: 1, fullName: 'John Doe', post: 'Developer'};
      EmployeeService.getProfile.mockResolvedValue(mockProfile);

      const profile = await employee.actions.getProfile({commit});

      expect(commit).toHaveBeenCalledWith('setFilled', true);
      expect(commit).toHaveBeenCalledWith('setProfile', mockProfile);
      expect(profile).toEqual(mockProfile);
    });

    it('getProfile action should handle error', async () => {
      const mockError = new Error('Get profile failed');
      EmployeeService.getProfile.mockRejectedValue(mockError);

      await expect(employee.actions.getProfile({commit})).rejects.toThrow(mockError);
    });

    it('getEmployeePhoto action should commit setPhotoURL on success', async () => {
      const mockImageURL = 'http://example.com/photo.png';
      EmployeeService.getEmployeePhoto.mockResolvedValue(mockImageURL);

      const imageURL = await employee.actions.getEmployeePhoto({commit});

      expect(commit).toHaveBeenCalledWith('setPhotoURL', mockImageURL);
      expect(imageURL).toEqual(mockImageURL);
    });

    it('getEmployeePhoto action should handle error', async () => {
      const mockError = new Error('Get employee photo failed');
      EmployeeService.getEmployeePhoto.mockRejectedValue(mockError);

      await expect(employee.actions.getEmployeePhoto({commit})).rejects.toThrow(mockError);
    });

    it('getEmployees action should return employees', async () => {
      const mockEmployees = [{id: 1, fullName: 'John Doe'}];
      const params = {searchQuery: '', searchBy: 'full_name', sortDirection: 'ASC'};
      EmployeeService.getEmployees.mockResolvedValue(mockEmployees);

      const employees = await employee.actions.getEmployees({commit}, params);

      expect(employees).toEqual(mockEmployees);
    });

    it('getEmployee action should return employee', async () => {
      const mockEmployee = {id: 1, fullName: 'John Doe'};
      const infoCardID = 1;
      EmployeeService.getEmployee.mockResolvedValue(mockEmployee);

      const employeeData = await employee.actions.getEmployee({commit}, infoCardID);

      expect(employeeData).toEqual(mockEmployee);
    });

    it('confirmEmployeeCard action should return response on success', async () => {
      const infoCardID = 1;
      const mockResponse = {message: 'Employee card confirmed'};
      EmployeeService.confirmEmployeeCard.mockResolvedValue(mockResponse);

      const response = await employee.actions.confirmEmployeeCard({commit}, infoCardID);

      expect(response).toEqual(mockResponse);
    });

    it('createEmployeePassage action should return response on success', async () => {
      const passageInfo = {employeeId: 1, access: true};
      const mockResponse = {message: 'Passage created'};
      EmployeeService.createEmployeePassage.mockResolvedValue(mockResponse);

      const response = await employee.actions.createEmployeePassage({commit}, passageInfo);

      expect(response).toEqual(mockResponse);
    });
  });

  describe('mutations', () => {
    it('setFilled mutation should update state correctly', () => {
      const state = {status: {filled: false}};

      employee.mutations.setFilled(state, true);

      expect(state.status.filled).toBe(true);
    });

    it('setProfile mutation should update state correctly', () => {
      const state = {profile: null, status: {filled: false}};
      const mockProfile = {id: 1, fullName: 'John Doe'};

      employee.mutations.setProfile(state, mockProfile);

      expect(state.profile).toEqual(mockProfile);
      expect(state.status.filled).toBe(true);
    });

    it('setPhotoURL mutation should update photoURL correctly', () => {
      const state = {photoURL: '//ssl.gstatic.com/accounts/ui/avatar_2x.png'};
      const newPhotoURL = 'http://example.com/newPhoto.png';

      employee.mutations.setPhotoURL(state, newPhotoURL);

      expect(state.photoURL).toBe(newPhotoURL);
    });
  });
});
