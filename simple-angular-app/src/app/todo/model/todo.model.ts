export default interface Todo {
  id: string | undefined;
  title: string | undefined;
  description: string | undefined;
  is_completed: boolean | undefined;
  createdAt: Date |  undefined;
  updatedAt: Date | undefined;
}
