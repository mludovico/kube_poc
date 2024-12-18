import {Component, Input, isDevMode} from '@angular/core';
import Todo from './model/todo.model';
import {TodoService} from './service/todo.service';

@Component({
  selector: 'app-todo',
  standalone: true,
  imports: [],
  templateUrl: './todo.component.html',
  styleUrl: './todo.component.css'
})
export class TodoComponent {
  constructor(private todoService: TodoService) {}

  @Input({required: true})
  todo!: Todo;

  complete() {
    this.todoService.completeTodo(this.todo).subscribe(() => {
      this.todo.is_completed = true;
      console.log('Todo completed');
    });
  }

  delete() {
    this.todoService.deleteTodo(this.todo).subscribe(()=> {
      console.log('Todo deleted');
    });
  }

  edit() {
    this.todoService.editTodo(this.todo).subscribe(todo => {
      this.todo = todo;
      console.log('Todo edited');
    });
  }

  protected readonly isDevMode = isDevMode;
  protected readonly navigator = navigator;
}
