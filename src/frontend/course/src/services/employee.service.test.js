import EmployeeService from './employee.service';
import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';

const mock = new MockAdapter(axios);
const API_URL = 'http://localhost:8081/api/v1/';

describe('employee.service', () => {
  beforeEach(() => {
    global.URL.createObjectURL = jest.fn(() => 'mocked-url');
  });

  afterEach(() => {
    mock.reset();
    localStorage.clear();
  });

  it('should fill profile and return response data', async () => {
    const formData = new FormData();
    formData.append('fullName', 'John Doe');

    const mockResponse = {success: true};
    mock.onPost(`${API_URL}profile`).reply(200, mockResponse);

    const data = await EmployeeService.fillProfile(formData);

    expect(data).toEqual(mockResponse);
  });

  it('should get profile data', async () => {
    const mockResponse = {fullName: 'John Doe', post: 'Developer'};

    mock.onGet(`${API_URL}profile`).reply(200, mockResponse);

    const data = await EmployeeService.getProfile();

    expect(data).toEqual(mockResponse);
  });

  it('should get employee photo and return blob URL', async () => {
    const mockPhoto = new ArrayBuffer(8);

    mock.onGet(`${API_URL}employee-photo`).reply(200, mockPhoto);

    const url = await EmployeeService.getEmployeePhoto();

    expect(url).toBeTruthy();
  });

  it('should get employees based on search query', async () => {
    const mockResponse = {infoCards: [{id: 1, fullName: 'John Doe'}]};

    mock.onGet(`${API_URL}infocards`).reply(200, mockResponse);

    const employees = await EmployeeService.getEmployees('John', 'full_name', 'ASC');

    expect(employees).toEqual(mockResponse.infoCards);
  });

  it('should get employee by id', async () => {
    const mockResponse = {id: 1, fullName: 'John Doe', post: 'Developer'};

    mock.onGet(`${API_URL}infocards/1`).reply(200, mockResponse);

    const employee = await EmployeeService.getEmployee(1);

    expect(employee).toEqual(mockResponse);
  });

  it('should get employee info card photo by id and return blob URL', async () => {
    const mockPhoto = new ArrayBuffer(8);

    mock.onGet(`${API_URL}infocard-photos/1`).reply(200, mockPhoto);

    const url = await EmployeeService.getEmployeeInfoCardPhoto(1);

    expect(url).toBeTruthy();
  });

  it('should confirm employee card by id', async () => {
    const mockResponse = {success: true};

    mock.onPatch(`${API_URL}infocards/1`).reply(200, mockResponse);

    const result = await EmployeeService.confirmEmployeeCard(1);

    expect(result).toEqual(mockResponse);
  });

  it('should create employee passage and return response data', async () => {
    const passageInfo = {infoCardID: 1, time: '2024-10-12T08:00:00Z'};

    const mockResponse = {success: true};

    mock.onPost(`${API_URL}passages`).reply(200, mockResponse);

    const result = await EmployeeService.createEmployeePassage(passageInfo);

    expect(result).toEqual(mockResponse);
  });
});
