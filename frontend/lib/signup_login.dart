import 'package:flutter/material.dart';

class AuthPage extends StatelessWidget {
  const AuthPage({super.key}); // constを追加

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 2, // タブの数
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Let's start!"), // constを追加
          backgroundColor: Theme.of(context).colorScheme.primaryContainer,
          bottom: const TabBar( // constを追加
            tabs: [
              Tab(text: "Login"), // constの適用は不要（既に不変）
              Tab(text: "SignUp"), // constの適用は不要（既に不変）
            ],
          ),
        ),
        body: const TabBarView( // constを追加
          children: [
            LoginForm(), // constを追加可能
            SignUpForm(), // constを追加可能
          ],
        ),
      ),
    );
  }
}

// Loginフォーム
class LoginForm extends StatelessWidget {
  const LoginForm({super.key}); // constを追加

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0), // constを追加
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const TextField(
            decoration: InputDecoration(
              labelText: 'Username',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16), // constを追加
          const TextField(
            decoration: InputDecoration(
              labelText: 'Password',
              border: OutlineInputBorder(),
            ),
            obscureText: true,
          ),
          const SizedBox(height: 20), // constを追加
          ElevatedButton(
            onPressed: () {
              print("Login pressed");
            },
            child: const Text("Login"), // constを追加
          ),
        ],
      ),
    );
  }
}

// SignUpフォーム
class SignUpForm extends StatelessWidget {
  const SignUpForm({super.key}); // constを追加

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0), // constを追加
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const TextField(
            decoration: InputDecoration(
              labelText: 'Username',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16), // constを追加
          const TextField(
            decoration: InputDecoration(
              labelText: 'Password',
              border: OutlineInputBorder(),
            ),
            obscureText: true,
          ),
          const SizedBox(height: 20), // constを追加
          ElevatedButton(
            onPressed: () {
              print("SignUp pressed");
            },
            child: const Text("SignUp"), // constを追加
          ),
        ],
      ),
    );
  }
}
