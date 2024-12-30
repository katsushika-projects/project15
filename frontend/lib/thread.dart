import 'package:flutter/material.dart';

class ThreadPage extends StatelessWidget {
  final String classId;
  final String className;

  const ThreadPage({super.key, required this.classId, required this.className});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Thread'),
        backgroundColor: Theme.of(context).colorScheme.primaryContainer,
      ),
      body: Text(classId)
    );
  }
}