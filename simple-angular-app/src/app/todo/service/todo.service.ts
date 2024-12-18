import {Injectable} from '@angular/core';
import Todo from '../model/todo.model';
import {Observable, shareReplay} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import * as api from '../../api'

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  constructor(private http: HttpClient) {
  }

  getTodos(): Observable<Todo[]> {
    // @ts-ignore
    return this.http.get(api.TODOS).pipe(shareReplay());
  }

  createTodo(todo: Todo): Observable<any> {
    return this.http.post(api.TODOS, todo).pipe(shareReplay());
  }

  completeTodo(todo: Todo): Observable<any> {
    return this.http.put(`${api.TODOS}/${todo.id}`, {
      ...todo,
      is_completed: true
    }).pipe(shareReplay());
  }

  editTodo(changes: Partial<Todo>): Observable<Todo> {
    return this.http.put(`${api.TODOS}/${changes.id}`, changes).pipe(
      // @ts-ignore
      shareReplay()
    );
  }

  deleteTodo(todo: Todo): Observable<any> {
    return this.http.delete(`${api.TODOS}/${todo.id}`).pipe(shareReplay());
  }
}
