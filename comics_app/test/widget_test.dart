import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:comics_app/main.dart';

void main() {
  testWidgets('Comics list smoke test', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(MyApp());

    // Verify that the comics list is initially empty.
    expect(find.text('Comics'), findsOneWidget);
    expect(find.byType(ListTile), findsNothing);

    // Wait for the comics to load.
    await tester.pumpAndSettle();

    // Verify that the comics list is populated.
    expect(find.byType(ListTile), findsWidgets);
  });
}
