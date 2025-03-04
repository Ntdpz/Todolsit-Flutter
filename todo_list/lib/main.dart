import 'package:flutter/material.dart';
import 'todolist/todolist.dart';
void main() {
  runApp(
     MaterialApp(
      home: Scaffold(
        body: Center(
          child: TodoList(),

        ),
      ),
    ),
  );
}
