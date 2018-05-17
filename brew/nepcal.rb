class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/nepcal/nepcal/archive/v0.1.1.tar.gz"
  head "https://github.com/nepcal/nepcal.git"
  version "0.1.1"
  sha256 "dd78f37bba0daad312fe83bafc1abf75480ba7db86047e1aa0244c8a151fa94a"
  depends_on "make" => :build

  def install
    bin.install "nepcal"
  end
end
