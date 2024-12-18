import {Component} from '@angular/core';
import Todo from '../../todo/model/todo.model';
import {FormsModule} from '@angular/forms';
import {TodoService} from '../../todo/service/todo.service';

@Component({
  selector: 'app-new-todo',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './new-todo.component.html',
  styleUrl: './new-todo.component.css'
})
export class NewTodoComponent {
  todo: Todo = {
    id: undefined,
    title: undefined,
    description: undefined,
    is_completed: false,
    createdAt: undefined,
    updatedAt: undefined
  };

  constructor(private todoService: TodoService) {
  }

  create() {
    this.todoService.createTodo(this.todo).subscribe(() => {
      console.log('Todo created');
    });
  }
}
