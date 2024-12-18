import {Component} from '@angular/core';
import {TodoService} from './todo/service/todo.service';
import Todo from './todo/model/todo.model';
import {TodoComponent} from './todo/todo.component';
import {NewTodoComponent} from './new_todo/new-todo/new-todo.component';
import {Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [TodoComponent, NewTodoComponent, AsyncPipe],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'simple-angular-app';
  todos$!: Observable<Todo[]>;
  todoService: TodoService;

  constructor(todoService: TodoService) {
    this.todoService = todoService;
    this.loadTodos();
  }

  loadTodos() {
    this.todos$ = this.todoService.getTodos();
    this.todos$.subscribe((todos) => {
      console.log(todos);
    });
  }
}
