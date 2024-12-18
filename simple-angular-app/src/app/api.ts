import { environment } from './environments/environment';

const API_URL = `${environment.apiHost}:${environment.apiPort}`;
export const TODOS = `${API_URL}/todos`;
