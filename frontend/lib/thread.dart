import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:flutter/material.dart';
import 'signup_login.dart';
import 'search.dart';

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
        automaticallyImplyLeading: false,
        leading: IconButton(
          icon: const Icon(Icons.home),
          onPressed: () {
            Navigator.pop(context);
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) => const SearchPage(),
                ),
              );
          },
          ),
      ),
      body: ThreadForm(classId: classId, className: className,),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          showDialog(
            context: context,
            builder: (context) => PostDialog(
              classId: classId,
              // 例: 投稿後にリストをリフレッシュしたい場合
              // onPost: () => _fetchThreads() のように、
              // ThreadPageからコールバックを渡すこともできます
              onPost: () {
                Navigator.pop(context); // ダイアログを閉じる
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('投稿が完了しました')),
                );
              },
              className: className,
            ),
          );
        },
        tooltip: '新しい投稿',
        child: const Icon(Icons.add),
      ),
    );
  }
}

class PostDialog extends StatelessWidget {
  final String classId;
  final VoidCallback onPost;
  final String className;

  const PostDialog({
    super.key,
    required this.classId,
    required this.onPost,
    required this.className,
  });

  @override
  Widget build(BuildContext context) {
    final TextEditingController controller = TextEditingController();

    return AlertDialog(
      title: const Text('新しい投稿'),
      content: TextField(
        controller: controller,
        decoration: const InputDecoration(hintText: '投稿を入力してください'),
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.pop(context);
          },
          child: const Text('キャンセル'),
        ),
        TextButton(
          onPressed: () async {
            final content = controller.text.trim();
            if (content.isEmpty) {
              if (context.mounted) {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('投稿内容を入力してください')),
                );
              }
              return;
            }
            try {
              final url = Uri.parse('http://localhost:8080/discription');
              final response = await http.post(
                url,
                headers: {
                  'Content-Type': 'application/json',
                  'Authorization': 'Bearer ${TokenManager().accessToken}',
                },
                body: jsonEncode({
                  'discript': content,
                  'class_id': classId,
                }),
              );
              if (response.statusCode == 201) {
                if (context.mounted) {
                  Navigator.pop(context);
                  Navigator.pushReplacement(
                    context,
                    MaterialPageRoute(
                      builder: (context) => ThreadPage(
                        classId: classId,
                        className: className,
                      ),
                    ),
                  );
                }
              } else {
                if (context.mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('投稿エラー: ${response.statusCode}')),
                  );
                }
              }
            } catch (e) {
              if (context.mounted) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('ネットワークエラー: $e')),
                );
              }
            }
          },
          child: const Text('投稿'),
        ),
      ],
    );
  }
}

class ThreadForm extends StatefulWidget {
  final String classId;
  final String className;

  const ThreadForm({super.key, required this.classId, required this.className});

  @override
  ThreadPageState createState() => ThreadPageState();
}

class ThreadPageState extends State<ThreadForm> {
  bool _isLoading = false;
  List<dynamic> _threads = [];

  @override
  void initState() {
    super.initState();
    _fetchThreads();
  }

  Future<void> _fetchThreads() async {
    setState(() {
      _isLoading = true;
    });

    final url = Uri.parse('http://localhost:8080/discription/get');
    try {
      final response = await http.post(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ${TokenManager().accessToken}',
        },
        body: jsonEncode({'class_id': widget.classId}),
      );

      if (response.statusCode == 200) {
        if (mounted) {
          final data = jsonDecode(response.body);
          setState(() {
            _threads = data['discripts'];
          });
        }
      } else {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Error: ${response.statusCode}')),
          );
        }
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Network error: $e')),
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          _isLoading = false;
        });
      }
    }
  }

  Future<void> _deleteThread(String threadId) async {
    setState(() {
      _isLoading = true;
    });

    final url = Uri.parse('http://localhost:8080/discription/$threadId');
    try {
      final response = await http.delete(
        url,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ${TokenManager().accessToken}',
        },
      );

      if (response.statusCode == 200) {
        if (mounted) {
          await showDialog(
            context: context,
            builder: (BuildContext context) {
              return SimpleDialog(
                title: const Text('本当に削除しますか？'),
                children: [
                  SimpleDialogOption(
                    onPressed: () {
                      Navigator.pop(context);
                      Navigator.pop(context);
                      Navigator.push(
                        context,
                        CupertinoPageRoute(
                          builder: (context) => ThreadPage(
                            classId: widget.classId,
                            className: widget.className,
                          ),
                        ),
                      );
                    },
                    child: const Text('削除'),
                  ),
                  SimpleDialogOption(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    child: const Text('キャンセル'),
                  )
                ],
              );
            }
            );
        }
      } else {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text('Server error: ${response.statusCode}')),
          );
        }
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Network error: $e'))
        );
      }
    } finally {
      if (mounted) {
        setState(() {
          _isLoading = false;
        });
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return const Center(
        child: CircularProgressIndicator(),
      );
    }

    return ListView.builder(
      itemCount: _threads.length,
      itemBuilder: (context, index) {
        final thread = _threads[index];
        return Container(
          margin: const EdgeInsets.symmetric(horizontal: 15, vertical: 8),
          padding: const EdgeInsets.all(8),
          decoration: BoxDecoration(
            color: Colors.white,
            border: Border.all(color: Colors.grey),
            borderRadius: BorderRadius.circular(8),
          ),
          child: Row(
            children: [
              Text(thread['Discript']),
              IconButton(
                onPressed: () async {
                  await _deleteThread(thread['ID']);
                },
                icon: const Icon(Icons.delete),
                )
            ],
          )
        );
      },
    );
  }
}
