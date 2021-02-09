//
//  EntryViewController.swift
//  CalorieMalorie
//
//  Created by Kamaal M Farah on 09/02/2021.
//

import UIKit
import SwiftUI

class EntryViewController: UIHostingController<EntryContentView> {

    override init(rootView: EntryContentView) {
        super.init(rootView: rootView)
    }

    @objc required dynamic init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)

        self.view.backgroundColor = .systemBackground
        self.title = "CalorieMalorie"
        self.navigationController?.navigationBar.prefersLargeTitles = true
    }

}

struct EntryContentView: View {
    var body: some View {
        Text("Hello world!")
    }
}
