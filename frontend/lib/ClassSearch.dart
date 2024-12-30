//import 'dart:nativewrappers/_internal/vm/lib/internal_patch.dart';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'signup_login.dart';
import 'thread.dart';

class ClassPage extends StatelessWidget {
  final String classId;

  const ClassPage({super.key, required this.classId});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Class"),
        backgroundColor: Theme.of(context).colorScheme.primaryContainer,
      ),
      body: ClassSearchPage(classId: classId),
    );
  }
}

class ClassSearchPage extends StatefulWidget {
  final String classId;

  const ClassSearchPage({super.key, required this.classId});

  @override
  ClassSearchPageState createState() => ClassSearchPageState();
}

class ClassSearchPageState extends State<ClassSearchPage> {
  /// 通信中かどうかを管理するフラグ
  bool _isLoading = false;

  /// 取得したクラス情報を保持する変数（null 許容型）
  Map<String, dynamic>? _classData;

  @override
  void initState() {
    super.initState();
    // ページ生成時にクラス情報を取得
    fetchClassById(widget.classId);
  }

  /// classId をもとに API からクラス情報を取得
  Future<void> fetchClassById(String classId) async {
    setState(() {
      _isLoading = true;
    });

    try {
      // 例: GET /api/classes/:classId
      // 実際のエンドポイントに合わせて変更してください
      final url = Uri.parse('http://localhost:8080/classes/$classId');
      final response = await http.get(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ${TokenManager().accessToken}',
        },
        );
      
      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);

        // ここでは data が { "id": 1, "name": "Math", "teacher": "John Doe" } のように
        // 1つのクラス情報を返すと仮定
        setState(() {
          _classData = data; 
        });
      } else {
        // エラーが返ってきた場合
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Error: ${response.statusCode}')),
          );
        }
      }
    } catch (e) {
      // ネットワークエラーや JSON パースエラーなど
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Error: $e')),
        );
      }
    } finally {
      // ローディング終了
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    // 通信中はローディングインジケーターを表示
    if (_isLoading) {
      return const Center(
        child: CircularProgressIndicator(),
      );
    }

    // データがまだ取得できていない（または null）の場合
    if (_classData == null) {
      return const Center(
        child: Text('No class data found.'),
      );
    }

    // ここまで来たら _classData に情報が入っている
    // 例: {"id": 1, "name": "Math", "teacher": "John Doe"}
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start, // 左揃え
        children: [
          Text(
            'Class Name: ${_classData!['class']['ClassName']}',
            style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
          ),
          const SizedBox(height: 8),
          Text('ID: ${_classData!['class']['ID']}'),
          const SizedBox(height: 16),
          _isLoading
              ? const CircularProgressIndicator()
              : ElevatedButton(
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => ThreadPage(
                        classId: _classData!['class']['ID'],
                        className: _classData!['class']['ClassName'],
                      ),
                    ),
                  );
                },
                child: const Text("スレッドに参加"),
              ),
        ],
      ),
    );
  }
}
